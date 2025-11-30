package models

// MaterialNode represents a node in the materials tree (can be folder or material)
type MaterialNode struct {
	ID           string          `json:"id"`
	Name         string          `json:"name"`
	Type         string          `json:"type"` // "folder" or "material"
	ParentID     string          `json:"parentId,omitempty"`
	Children     []*MaterialNode `json:"children,omitempty"`
	MaterialType string          `json:"materialType,omitempty"` // PDF, Vídeo, Link, Documento
	URL          string          `json:"url,omitempty"`
	Description  string          `json:"description,omitempty"`
	DateAdded    string          `json:"dateAdded,omitempty"`
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
	URL          string `json:"url" binding:"required"`
	Description  string `json:"description"`
}

// UpdateNodeRequest represents the request to update a node
type UpdateNodeRequest struct {
	Name         string `json:"name,omitempty"`
	MaterialType string `json:"materialType,omitempty"`
	URL          string `json:"url,omitempty"`
	Description  string `json:"description,omitempty"`
}

// MoveNodeRequest represents the request to move a node
type MoveNodeRequest struct {
	NewParentID string `json:"newParentId" binding:"required"`
}
