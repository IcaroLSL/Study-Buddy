package handlers

import (
	"net/http"
	"strconv"
	"studybuddy/models"
	"studybuddy/storage"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateSubjectRequest representa o corpo da requisição para criar uma matéria
type CreateSubjectRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	ParentID    *int64 `json:"parentId"`
	Color       string `json:"color"`
	Icon        string `json:"icon"`
}

// UpdateSubjectRequest representa o corpo da requisição para atualizar uma matéria
type UpdateSubjectRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Color       string `json:"color"`
	Icon        string `json:"icon"`
}

// MoveSubjectRequest representa o corpo da requisição para mover uma matéria
type MoveSubjectRequest struct {
	NewParentID *int64 `json:"newParentId"`
}

// HandleGetSubjects retorna todas as matérias em formato de árvore
func HandleGetSubjects(c *gin.Context) {
	data, err := storage.LoadData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao carregar dados"})
		return
	}

	c.JSON(http.StatusOK, data.Subjects)
}

// HandleCreateSubject cria uma nova matéria (raiz ou sub-matéria)
func HandleCreateSubject(c *gin.Context) {
	var req CreateSubjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	data, err := storage.LoadData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao carregar dados"})
		return
	}

	newSubject := models.Subject{
		ID:          storage.GetNextSubjectID(),
		Name:        req.Name,
		Description: req.Description,
		ParentID:    req.ParentID,
		Color:       req.Color,
		Icon:        req.Icon,
		CreatedAt:   time.Now().Format(time.RFC3339),
		Children:    []models.Subject{},
	}

	// Se parentID foi fornecido, verificar se o pai existe
	if req.ParentID != nil {
		parent := storage.FindSubjectByID(data.Subjects, *req.ParentID)
		if parent == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Matéria pai não encontrada"})
			return
		}
	}

	if !storage.AddSubject(&data.Subjects, newSubject, req.ParentID) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao adicionar matéria"})
		return
	}

	if err := storage.SaveData(data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar dados"})
		return
	}

	c.JSON(http.StatusCreated, newSubject)
}

// HandleUpdateSubject atualiza uma matéria existente
func HandleUpdateSubject(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req UpdateSubjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	data, err := storage.LoadData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao carregar dados"})
		return
	}

	existing := storage.FindSubjectByID(data.Subjects, id)
	if existing == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Matéria não encontrada"})
		return
	}

	updated := models.Subject{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		ParentID:    existing.ParentID,
		Color:       req.Color,
		Icon:        req.Icon,
		CreatedAt:   existing.CreatedAt,
	}

	if !storage.UpdateSubjectInTree(&data.Subjects, updated) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar matéria"})
		return
	}

	if err := storage.SaveData(data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar dados"})
		return
	}

	// Retornar a matéria atualizada com os filhos
	updatedSubject := storage.FindSubjectByID(data.Subjects, id)
	c.JSON(http.StatusOK, updatedSubject)
}

// HandleDeleteSubject remove uma matéria (e seus filhos)
func HandleDeleteSubject(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	data, err := storage.LoadData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao carregar dados"})
		return
	}

	if !storage.RemoveSubject(&data.Subjects, id) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Matéria não encontrada"})
		return
	}

	if err := storage.SaveData(data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar dados"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "matéria removida"})
}

// HandleMoveSubject move uma matéria para outro pai
func HandleMoveSubject(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req MoveSubjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	data, err := storage.LoadData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao carregar dados"})
		return
	}

	// Verificar se a matéria existe
	existing := storage.FindSubjectByID(data.Subjects, id)
	if existing == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Matéria não encontrada"})
		return
	}

	// Se newParentID foi fornecido, verificar se o pai existe
	if req.NewParentID != nil {
		parent := storage.FindSubjectByID(data.Subjects, *req.NewParentID)
		if parent == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Novo pai não encontrado"})
			return
		}

		// Verificar se não está tentando mover para si mesmo ou para um descendente
		if *req.NewParentID == id {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Não é possível mover uma matéria para si mesma"})
			return
		}

		// Verificar se o novo pai não é um descendente
		if isDescendant(existing.Children, *req.NewParentID) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Não é possível mover uma matéria para um de seus descendentes"})
			return
		}
	}

	if !storage.MoveSubject(&data.Subjects, id, req.NewParentID) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao mover matéria"})
		return
	}

	if err := storage.SaveData(data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar dados"})
		return
	}

	// Retornar a matéria movida
	movedSubject := storage.FindSubjectByID(data.Subjects, id)
	c.JSON(http.StatusOK, movedSubject)
}

// isDescendant verifica se um ID é descendente na árvore de filhos
func isDescendant(children []models.Subject, id int64) bool {
	for _, child := range children {
		if child.ID == id {
			return true
		}
		if isDescendant(child.Children, id) {
			return true
		}
	}
	return false
}
