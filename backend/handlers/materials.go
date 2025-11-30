package handlers

import (
	"net/http"
	"studybuddy/models"
	"studybuddy/storage"

	"github.com/gin-gonic/gin"
)

// HandleGetMaterials returns the entire materials tree
func HandleGetMaterials(c *gin.Context) {
	tree, err := storage.LoadMaterials()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao carregar materiais"})
		return
	}
	c.JSON(http.StatusOK, tree)
}

// HandleGetMaterialNode returns a specific node and its children
func HandleGetMaterialNode(c *gin.Context) {
	nodeID := c.Param("id")

	tree, err := storage.LoadMaterials()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao carregar materiais"})
		return
	}

	node := storage.FindNodeByID(tree.Root, nodeID)
	if node == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Nó não encontrado"})
		return
	}

	c.JSON(http.StatusOK, node)
}

// HandleCreateFolder creates a new folder
func HandleCreateFolder(c *gin.Context) {
	var req models.CreateFolderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	tree, err := storage.LoadMaterials()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao carregar materiais"})
		return
	}

	folder, err := storage.AddFolder(tree, req.Name, req.ParentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, folder)
}

// HandleCreateMaterial creates a new material
func HandleCreateMaterial(c *gin.Context) {
	var req models.CreateMaterialRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	tree, err := storage.LoadMaterials()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao carregar materiais"})
		return
	}

	material, err := storage.AddMaterial(tree, req.Name, req.ParentID, req.MaterialType, req.URL, req.Description)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, material)
}

// HandleUpdateNode updates a node's properties
func HandleUpdateNode(c *gin.Context) {
	nodeID := c.Param("id")

	var req models.UpdateNodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	tree, err := storage.LoadMaterials()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao carregar materiais"})
		return
	}

	if err := storage.UpdateNode(tree, nodeID, req.Name, req.MaterialType, req.URL, req.Description); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return the updated node
	node := storage.FindNodeByID(tree.Root, nodeID)
	c.JSON(http.StatusOK, node)
}

// HandleDeleteNode deletes a node and its children
func HandleDeleteNode(c *gin.Context) {
	nodeID := c.Param("id")

	tree, err := storage.LoadMaterials()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao carregar materiais"})
		return
	}

	if err := storage.DeleteNode(tree, nodeID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "removido"})
}

// HandleMoveNode moves a node to a different parent
func HandleMoveNode(c *gin.Context) {
	nodeID := c.Param("id")

	var req models.MoveNodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	tree, err := storage.LoadMaterials()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao carregar materiais"})
		return
	}

	if err := storage.MoveNode(tree, nodeID, req.NewParentID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return the updated tree
	updatedTree, err := storage.LoadMaterials()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao carregar materiais atualizados"})
		return
	}
	c.JSON(http.StatusOK, updatedTree)
}
