package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
    Notes     []Note     `json:"notes"`
    Reminders []Reminder `json:"reminders"`
    StudyLog  map[string]int `json:"studyLog"`
    Subjects  []string   `json:"subjects"`
    Events    []Event    `json:"events"` // <-- ADICIONE ISTO
}


var (
	dataFile = "data.json"
	mutex    sync.Mutex
)

func main() {
	r := gin.Default()

	// Permitir CORS para qualquer origem
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // ou especifique seu frontend
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: false,
	}))

	r.GET("/data", func(c *gin.Context) {
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
	})

	r.POST("/data", func(c *gin.Context) {
		mutex.Lock()
		defer mutex.Unlock()
	
		var data AppData
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
			return
		}
	
		// Debug: verifique se os eventos estão sendo recebidos
		fmt.Printf("Eventos recebidos: %d\n", len(data.Events))
		for _, event := range data.Events {
			fmt.Printf("Evento: %s em %s às %s\n", event.Title, event.Date, event.Time)
		}
	
		bytes, _ := json.MarshalIndent(data, "", "  ")
		_ = ioutil.WriteFile(dataFile, bytes, 0644)
		c.JSON(http.StatusOK, gin.H{"status": "salvo"})
	})

	r.Run(":8080")
}
