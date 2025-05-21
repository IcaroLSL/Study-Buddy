package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"github.com/gin-gonic/gin"
    "time"
)

type MyResumes struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date    string `json:"date"`
}

var dataSet []MyResumes
const dataFile = "data_resumes.json"

func main() {
	loadData(dataFile)
	printData()

	router := gin.Default()

	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Use(corsMiddleware())
	
	router.GET("/", homeHandler)
	router.GET("/ResumesData", getCardsHandler)
	router.GET("/ResumesData/get", getCardByIDHandler)
	router.POST("/ResumesData/add", postCardHandler)
	router.PUT("/ResumesData/update", updateHandler)
	router.DELETE("/ResumesData/dell", deleteHandler)

	fmt.Println("\nServidor rodando em http://localhost:8080")
	router.Run(":8080")
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	}
}

func loadData(filename string) {
	file, err := os.ReadFile(filename)
	if err != nil {log.Fatal("Error to read file:", err)}
	if err := json.Unmarshal(file, &dataSet); err != nil {log.Fatal("Error to decode JSON:", err)}
}

func saveData() {
	file, err := json.MarshalIndent(dataSet, "", "  ")
	if err != nil {log.Fatal("Error to encoder JSON:", err)}
	if err := os.WriteFile(dataFile, file, 0644); err != nil {log.Fatal("Error to save file:", err)}
}

func printData() {
	fmt.Println("\nData Load:")
	for _, card := range dataSet {
		fmt.Printf("\nID : %d\nTítulo : %s\nDescrição : %s\nData : %s\n", card.Id, card.Title, card.Description, card.Date,)
	}
}

func homeHandler(c *gin.Context) {
	c.String(http.StatusOK, "Study-Buddy\nEndpoints:\n-GET : /ResumesData\n-GET : /ResumesData/get?id=1\n-POST : /ResumesData/add\n-PUT : /ResumesData/update?id=1\n-DELETE : /ResumesData/dell?id=1")
}

func getCardsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, dataSet)
}

func getCardByIDHandler(c *gin.Context) {
    idStr := c.Query("id")
    if idStr == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Parameter 'id' is required",
        })
        return
    }
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid ID",
            "details": err.Error(),
        })
        return
    }
    for _, card := range dataSet {
        if card.Id == id {
            c.JSON(http.StatusOK, card)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{
        "error": fmt.Sprintf("Card with ID %d not found", id),
    })
}

func postCardHandler(c *gin.Context) {
	var newResume MyResumes
	if err := c.ShouldBindJSON(&newResume); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	newResume.Id = len(dataSet) + 1
	newResume.Date = formatarDataPortugues(time.Now())
    dataSet = append(dataSet, newResume)
	saveData()
	c.JSON(http.StatusCreated, newResume)
}

func updateHandler(c *gin.Context) {
    idStr := c.Query("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid ID",
            "details": err.Error(),
        })
        return
    }
    var updatedResume MyResumes
    if err := c.ShouldBindJSON(&updatedResume); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid data",
            "details": err.Error(),
        })
        return
    }
    for i, card := range dataSet {
        if card.Id == id {
            updatedResume.Id = id
            updatedResume.Date = formatarDataPortugues(time.Now())
            dataSet[i] = updatedResume
            saveData()
            c.JSON(http.StatusOK, gin.H{
                "message": "Card updated successfully",
                "card": updatedResume,
            })
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{
        "error": fmt.Sprintf("Card with ID %d not found", id),
    })
}

func deleteHandler(c *gin.Context) {
    idStr := c.Query("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid ID",
            "details": err.Error(),
        })
        return
    }
    for i, card := range dataSet {
        if card.Id == id {
            dataSet = append(dataSet[:i], dataSet[i+1:]...)
            saveData()
            c.JSON(http.StatusOK, gin.H{
                "message": fmt.Sprintf("Card %d remove successfully", id),
            })
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{
        "error": fmt.Sprintf("Card with ID %d not found", id),
    })
}

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	dataAtual := time.Now()
// 	dataPT := formatarDataPortugues(dataAtual)
// 	fmt.Println("Data em português:", dataPT)
// }

func formatarDataPortugues(t time.Time) string {
	mesesPT := map[time.Month]string{
		time.January:   "Jan",
		time.February:  "Fev",
		time.March:     "Mar",
		time.April:     "Abr",
		time.May:      "Mai",
		time.June:      "Jun",
		time.July:      "Jul",
		time.August:    "Ago",
		time.September: "Set",
		time.October:   "Out",
		time.November:  "Nov",
		time.December:  "Dez",
	}
	
	return fmt.Sprintf("%02d %s %d", t.Day(), mesesPT[t.Month()], t.Year())
}