package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Reminder struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Date      string `json:"date"`
	Time      string `json:"time"`
	Priority  string `json:"priority"`
	Completed bool   `json:"completed"`
}

type Note struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Subject string `json:"subject"`
	Date    string `json:"date"`
}

type Event struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Date        string `json:"date"`
	Time        string `json:"time"`
	Description string `json:"description"`
}

type AppData struct {
	Notes     []Note         `json:"notes"`
	Reminders []Reminder     `json:"reminders"`
	StudyLog  map[string]int `json:"studyLog"`
	Subjects  []string       `json:"subjects"`
	Events    []Event        `json:"events"`
	Materials []Material     `json:"materials"`
	Folders   []string       `json:"folders"`
}

type Material struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Folder      string `json:"folder"`
	URL         string `json:"url"`
	DateAdded   string `json:"dateAdded"`
}

var (
	dataFile  = "data.json"
	mutex     sync.Mutex
	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:5500"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// protected := r.Group("/")
	// protected.Use(authMiddleware())
	// {
	// 	protected.GET("/data", handleGetData)
	// 	protected.POST("/data", handleSaveData)
	// 	protected.DELETE("/events/:id", handleDeleteEvent)
	// }
	// Desativando temporariamente a autenticação
	r.GET("/data", handleGetData)
	r.POST("/data", handleSaveData)
	r.DELETE("/events/:id", handleDeleteEvent)

	r.Run(":8080")
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token de autenticação não fornecido"})
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("método de assinatura inválido")
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}

		c.Next()
	}
}

func handleGetData(c *gin.Context) {
	mutex.Lock()
	defer mutex.Unlock()

	bytes, err := ioutil.ReadFile(dataFile)
	if err != nil {
		c.JSON(http.StatusOK, AppData{})
		return
	}

	var data AppData
	json.Unmarshal(bytes, &data)
	c.JSON(http.StatusOK, data)
}

func handleSaveData(c *gin.Context) {
	mutex.Lock()
	defer mutex.Unlock()

	var data AppData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	bytes, _ := json.MarshalIndent(data, "", "  ")
	_ = ioutil.WriteFile(dataFile, bytes, 0644)
	c.JSON(http.StatusOK, gin.H{"status": "salvo"})
}

func handleDeleteEvent(c *gin.Context) {
	mutex.Lock()
	defer mutex.Unlock()

	bytes, err := ioutil.ReadFile(dataFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao ler dados"})
		return
	}

	var data AppData
	if err := json.Unmarshal(bytes, &data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar dados"})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	newEvents := make([]Event, 0)
	for _, event := range data.Events {
		if event.ID != id {
			newEvents = append(newEvents, event)
		}
	}

	data.Events = newEvents

	updatedBytes, _ := json.MarshalIndent(data, "", "  ")
	_ = ioutil.WriteFile(dataFile, updatedBytes, 0644)
	c.JSON(http.StatusOK, gin.H{"status": "evento removido"})
}
