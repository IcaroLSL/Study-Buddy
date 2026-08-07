package models

// MaterialNode represents a node in the materials tree (can be folder or material)
type MaterialNode struct {
	ID           string          `json:"id"`
	Name         string          `json:"name"`
	Type         string          `json:"type"` // "folder" or "material"
	ParentID     string          `json:"parentId,omitempty"`
	Children     []*MaterialNode `json:"children,omitempty"`
	MaterialType string          `json:"materialType,omitempty"` // PDF, Vídeo, Link, Documento, Imagem
	URL          string          `json:"url,omitempty"`          // For external links
	FilePath     string          `json:"filePath,omitempty"`     // Local file path
	FileName     string          `json:"fileName,omitempty"`     // Original file name
	FileSize     int64           `json:"fileSize,omitempty"`     // File size in bytes
	Description  string          `json:"description,omitempty"`
	DateAdded    string          `json:"dateAdded,omitempty"`
	IsFile       bool            `json:"isFile"` // true = local file, false = external link
}

// MaterialsTree represents the root structure for materials
type MaterialsTree struct {
	Root *MaterialNode `json:"root"`
}

// CreateFolderRequest represents the request to create a new folder
type CreateFolderRequest struct {
	Name     string `json:"name" binding:"required"`
	ParentID string `json:"parentId" binding:"required"`
}

// CreateMaterialRequest represents the request to create a new material
type CreateMaterialRequest struct {
	Name         string `json:"name" binding:"required"`
	ParentID     string `json:"parentId" binding:"required"`
	MaterialType string `json:"materialType" binding:"required"`
	URL          string `json:"url"`                        // For links (optional when isFile is true)
	FilePath     string `json:"filePath,omitempty"`         // For uploaded files
	FileName     string `json:"fileName,omitempty"`         // Original file name
	FileSize     int64  `json:"fileSize,omitempty"`         // File size in bytes
	Description  string `json:"description"`
	IsFile       bool   `json:"isFile"` // true = file, false = link
}

// UpdateNodeRequest represents the request to update a node
type UpdateNodeRequest struct {
	Name         string `json:"name,omitempty"`
	MaterialType string `json:"materialType,omitempty"`
	URL          string `json:"url,omitempty"`
	FilePath     string `json:"filePath,omitempty"`
	FileName     string `json:"fileName,omitempty"`
	FileSize     int64  `json:"fileSize,omitempty"`
	Description  string `json:"description,omitempty"`
	IsFile       bool   `json:"isFile,omitempty"`
}

// MoveNodeRequest represents the request to move a node
type MoveNodeRequest struct {
	NewParentID string `json:"newParentId" binding:"required"`
}
