package handlers

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"studybuddy/models"
	"studybuddy/storage"
	"time"

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

	// Validate: either URL or file info must be provided
	if !req.IsFile && req.URL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL é obrigatória para materiais do tipo link"})
		return
	}

	if req.IsFile && req.FilePath == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "FilePath é obrigatório para materiais do tipo arquivo"})
		return
	}

	tree, err := storage.LoadMaterials()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao carregar materiais"})
		return
	}

	material, err := storage.AddMaterialWithFile(tree, req.Name, req.ParentID, req.MaterialType, req.URL, req.FilePath, req.FileName, req.FileSize, req.Description, req.IsFile)
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

// Constants for file upload
const (
	maxFileSize = 50 << 20 // 50 MB
	uploadsDir  = "storage/uploads"
)

// Allowed file extensions and their MIME types
var allowedExtensions = map[string][]string{
	".pdf":  {"application/pdf"},
	".doc":  {"application/msword"},
	".docx": {"application/vnd.openxmlformats-officedocument.wordprocessingml.document"},
	".ppt":  {"application/vnd.ms-powerpoint"},
	".pptx": {"application/vnd.openxmlformats-officedocument.presentationml.presentation"},
	".xls":  {"application/vnd.ms-excel"},
	".xlsx": {"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"},
	".txt":  {"text/plain"},
	".jpg":  {"image/jpeg"},
	".jpeg": {"image/jpeg"},
	".png":  {"image/png"},
	".gif":  {"image/gif"},
	".mp4":  {"video/mp4"},
	".mp3":  {"audio/mpeg"},
}

// isAllowedExtension checks if file extension is allowed
func isAllowedExtension(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	_, exists := allowedExtensions[ext]
	return exists
}

// sanitizeFileName removes potentially dangerous characters from filename
func sanitizeFileName(filename string) string {
	// Remove path separators and other dangerous characters
	filename = filepath.Base(filename)
	// Replace any remaining path separators
	filename = strings.ReplaceAll(filename, "/", "_")
	filename = strings.ReplaceAll(filename, "\\", "_")
	filename = strings.ReplaceAll(filename, "..", "_")
	return filename
}

// generateUniqueFileName generates a unique filename to avoid conflicts
func generateUniqueFileName(originalName string) string {
	ext := filepath.Ext(originalName)
	nameWithoutExt := strings.TrimSuffix(originalName, ext)
	timestamp := time.Now().UnixNano()
	return fmt.Sprintf("%s_%d%s", sanitizeFileName(nameWithoutExt), timestamp, ext)
}

// getContentType returns the content type for a file extension
func getContentType(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".pdf":
		return "application/pdf"
	case ".doc":
		return "application/msword"
	case ".docx":
		return "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	case ".ppt":
		return "application/vnd.ms-powerpoint"
	case ".pptx":
		return "application/vnd.openxmlformats-officedocument.presentationml.presentation"
	case ".xls":
		return "application/vnd.ms-excel"
	case ".xlsx":
		return "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	case ".txt":
		return "text/plain"
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	case ".mp4":
		return "video/mp4"
	case ".mp3":
		return "audio/mpeg"
	default:
		return "application/octet-stream"
	}
}

// HandleUploadFile handles file upload
func HandleUploadFile(c *gin.Context) {
	// Parse multipart form with size limit
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxFileSize)

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		if err.Error() == "http: request body too large" {
			c.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": "Arquivo muito grande. Limite: 50MB"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao receber arquivo"})
		return
	}
	defer func(file multipart.File) {
		_ = file.Close()
	}(file)

	// Validate file extension
	if !isAllowedExtension(header.Filename) {
		c.JSON(http.StatusUnsupportedMediaType, gin.H{
			"error": "Tipo de arquivo não permitido. Tipos permitidos: PDF, DOC, DOCX, PPT, PPTX, XLS, XLSX, TXT, JPG, JPEG, PNG, GIF, MP4, MP3",
		})
		return
	}

	// Ensure uploads directory exists
	if err := os.MkdirAll(uploadsDir, 0755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar diretório de uploads"})
		return
	}

	// Generate unique filename
	uniqueFileName := generateUniqueFileName(header.Filename)
	filePath := filepath.Join(uploadsDir, uniqueFileName)

	// Create the file
	dst, err := os.Create(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar arquivo"})
		return
	}
	defer func(dst *os.File) {
		_ = dst.Close()
	}(dst)

	// Copy file content
	written, err := io.Copy(dst, file)
	if err != nil {
		// Clean up partial file
		_ = os.Remove(filePath)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar arquivo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"filePath": filePath,
		"fileName": header.Filename,
		"fileSize": written,
	})
}

// HandleDownloadFile handles file download
func HandleDownloadFile(c *gin.Context) {
	materialID := c.Param("id")

	tree, err := storage.LoadMaterials()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao carregar materiais"})
		return
	}

	node := storage.FindNodeByID(tree.Root, materialID)
	if node == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Material não encontrado"})
		return
	}

	if node.Type != "material" || !node.IsFile {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Este material não é um arquivo"})
		return
	}

	// Validate file path to prevent path traversal
	cleanPath := filepath.Clean(node.FilePath)
	if !strings.HasPrefix(cleanPath, uploadsDir) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Acesso não autorizado"})
		return
	}

	// Check if file exists
	if _, err := os.Stat(cleanPath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Arquivo não encontrado no servidor"})
		return
	}

	// Set headers for download
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", node.FileName))
	c.Header("Content-Type", getContentType(node.FileName))
	c.File(cleanPath)
}

// HandleViewFile handles inline file viewing
func HandleViewFile(c *gin.Context) {
	materialID := c.Param("id")

	tree, err := storage.LoadMaterials()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao carregar materiais"})
		return
	}

	node := storage.FindNodeByID(tree.Root, materialID)
	if node == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Material não encontrado"})
		return
	}

	if node.Type != "material" || !node.IsFile {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Este material não é um arquivo"})
		return
	}

	// Validate file path to prevent path traversal
	cleanPath := filepath.Clean(node.FilePath)
	if !strings.HasPrefix(cleanPath, uploadsDir) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Acesso não autorizado"})
		return
	}

	// Check if file exists
	if _, err := os.Stat(cleanPath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Arquivo não encontrado no servidor"})
		return
	}

	// Set headers for inline viewing
	contentType := getContentType(node.FileName)
	c.Header("Content-Disposition", fmt.Sprintf("inline; filename=\"%s\"", node.FileName))
	c.Header("Content-Type", contentType)
	c.File(cleanPath)
}
