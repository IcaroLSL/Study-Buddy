package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strconv"
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

type Material struct {
    ID          int64  `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Type        string `json:"type"`
    Folder      string `json:"folder"`
    URL         string `json:"url"`
    DateAdded   string `json:"dateAdded"`
}

type AppData struct {
    Notes     []Note     `json:"notes"`
    Reminders []Reminder `json:"reminders"`
    StudyLog  map[string]int `json:"studyLog"`
    Subjects  []string   `json:"subjects"`
    Events    []Event    `json:"events"`
    Materials []Material `json:"materials"`
    Folders   []string   `json:"folders"`
}

var (
	dataFile = "data.json"
	mutex    sync.Mutex
)

func main() {
	r := gin.Default()

	// Permitir CORS para qualquer origem
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
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

	r.DELETE("/events/:id", func(c *gin.Context) {
		mutex.Lock()
		defer mutex.Unlock()
	
		// Ler o arquivo atual
		bytes, err := ioutil.ReadFile(dataFile)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao ler dados"})
			return
		}
	
		var data AppData
		if err := json.Unmarshal(bytes, &data); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao decodificar dados"})
			return
		}
		
		// Converter ID de string para int64
		idStr := c.Param("id")
		fmt.Printf("Recebida requisição para deletar evento ID: %s\n", idStr)
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}
	
		// Filtrar o evento a ser removido
		newEvents := make([]Event, 0)
		for _, event := range data.Events {
			if event.ID != id {
				newEvents = append(newEvents, event)
			}
		}
	
		// Atualizar os eventos
		data.Events = newEvents
	
		// Salvar de volta no arquivo
		updatedBytes, _ := json.MarshalIndent(data, "", "  ")
		if err := ioutil.WriteFile(dataFile, updatedBytes, 0644); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar dados"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "evento removido"})
	})

	r.OPTIONS("/*cors", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type")
		c.Status(http.StatusOK)
	})
	
	r.Run(":8080")
}
