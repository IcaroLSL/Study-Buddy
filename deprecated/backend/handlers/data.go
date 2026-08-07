package handlers

import (
	"net/http"
	"strconv"
	"studybuddy/models"
	"studybuddy/storage"

	"github.com/gin-gonic/gin"
)

// HandleGetData returns all application data
func HandleGetData(c *gin.Context) {
	data, err := storage.LoadData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao ler dados"})
		return
	}
	c.JSON(http.StatusOK, data)
}

// HandleSaveData saves all application data
func HandleSaveData(c *gin.Context) {
	var data models.AppData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	if err := storage.SaveData(data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar dados"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "salvo"})
}

// HandleDeleteEvent deletes a specific event by ID
func HandleDeleteEvent(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	_, err = storage.DeleteEvent(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar evento"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "evento removido"})
}
