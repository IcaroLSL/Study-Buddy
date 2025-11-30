package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"studybuddy/models"
	"sync"
	"time"
)

var (
	materialsFile  = "storage/materials.json"
	materialsMutex sync.Mutex
)

// LoadMaterials loads the materials tree from the JSON file
func LoadMaterials() (*models.MaterialsTree, error) {
	materialsMutex.Lock()
	defer materialsMutex.Unlock()

	bytes, err := os.ReadFile(materialsFile)
	if err != nil {
		if os.IsNotExist(err) {
			// Return default tree if file doesn't exist
			return createDefaultTree(), nil
		}
		return nil, err
	}

	var tree models.MaterialsTree
	if err := json.Unmarshal(bytes, &tree); err != nil {
		return nil, err
	}
	return &tree, nil
}

// createDefaultTree creates a default empty materials tree
func createDefaultTree() *models.MaterialsTree {
	return &models.MaterialsTree{
		Root: &models.MaterialNode{
			ID:       "root",
			Name:     "Raiz",
			Type:     "folder",
			Children: []*models.MaterialNode{},
		},
	}
}

// SaveMaterials saves the materials tree to the JSON file
func SaveMaterials(tree *models.MaterialsTree) error {
	materialsMutex.Lock()
	defer materialsMutex.Unlock()

	bytes, err := json.MarshalIndent(tree, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(materialsFile, bytes, 0644)
}

// FindNodeByID finds a node in the tree by its ID
func FindNodeByID(node *models.MaterialNode, id string) *models.MaterialNode {
	if node == nil {
		return nil
	}
	if node.ID == id {
		return node
	}
	for _, child := range node.Children {
		if found := FindNodeByID(child, id); found != nil {
			return found
		}
	}
	return nil
}

// FindParentOfNode finds the parent node of a given node ID
func FindParentOfNode(root *models.MaterialNode, targetID string) *models.MaterialNode {
	if root == nil {
		return nil
	}
	for _, child := range root.Children {
		if child.ID == targetID {
			return root
		}
		if found := FindParentOfNode(child, targetID); found != nil {
			return found
		}
	}
	return nil
}

// generateID generates a unique ID based on timestamp
func generateID(prefix string) string {
	return fmt.Sprintf("%s-%d", prefix, time.Now().UnixNano())
}

// AddFolder adds a new folder to the tree
func AddFolder(tree *models.MaterialsTree, name string, parentID string) (*models.MaterialNode, error) {
	parent := FindNodeByID(tree.Root, parentID)
	if parent == nil {
		return nil, errors.New("pasta pai não encontrada")
	}
	if parent.Type != "folder" {
		return nil, errors.New("o pai deve ser uma pasta")
	}

	newFolder := &models.MaterialNode{
		ID:       generateID("folder"),
		Name:     name,
		Type:     "folder",
		ParentID: parentID,
		Children: []*models.MaterialNode{},
	}

	parent.Children = append(parent.Children, newFolder)

	if err := SaveMaterials(tree); err != nil {
		return nil, err
	}
	return newFolder, nil
}

// AddMaterial adds a new material to the tree
func AddMaterial(tree *models.MaterialsTree, name, parentID, materialType, url, description string) (*models.MaterialNode, error) {
	parent := FindNodeByID(tree.Root, parentID)
	if parent == nil {
		return nil, errors.New("pasta pai não encontrada")
	}
	if parent.Type != "folder" {
		return nil, errors.New("o pai deve ser uma pasta")
	}

	newMaterial := &models.MaterialNode{
		ID:           generateID("material"),
		Name:         name,
		Type:         "material",
		ParentID:     parentID,
		MaterialType: materialType,
		URL:          url,
		Description:  description,
		DateAdded:    time.Now().Format("2006-01-02"),
	}

	parent.Children = append(parent.Children, newMaterial)

	if err := SaveMaterials(tree); err != nil {
		return nil, err
	}
	return newMaterial, nil
}

// DeleteNode removes a node from the tree
func DeleteNode(tree *models.MaterialsTree, nodeID string) error {
	if nodeID == "root" {
		return errors.New("não é possível excluir a pasta raiz")
	}

	parent := FindParentOfNode(tree.Root, nodeID)
	if parent == nil {
		return errors.New("nó não encontrado")
	}

	// Remove the node from parent's children
	newChildren := []*models.MaterialNode{}
	for _, child := range parent.Children {
		if child.ID != nodeID {
			newChildren = append(newChildren, child)
		}
	}
	parent.Children = newChildren

	return SaveMaterials(tree)
}

// MoveNode moves a node to a new parent
func MoveNode(tree *models.MaterialsTree, nodeID string, newParentID string) error {
	if nodeID == "root" {
		return errors.New("não é possível mover a pasta raiz")
	}

	// Find the node to move
	node := FindNodeByID(tree.Root, nodeID)
	if node == nil {
		return errors.New("nó não encontrado")
	}

	// Find the new parent
	newParent := FindNodeByID(tree.Root, newParentID)
	if newParent == nil {
		return errors.New("nova pasta pai não encontrada")
	}
	if newParent.Type != "folder" {
		return errors.New("o destino deve ser uma pasta")
	}

	// Check if trying to move a folder into itself or a descendant
	if node.Type == "folder" && (nodeID == newParentID || FindNodeByID(node, newParentID) != nil) {
		return errors.New("não é possível mover uma pasta para dentro de si mesma")
	}

	// Remove from current parent
	currentParent := FindParentOfNode(tree.Root, nodeID)
	if currentParent == nil {
		return errors.New("pai atual não encontrado")
	}

	newChildren := []*models.MaterialNode{}
	for _, child := range currentParent.Children {
		if child.ID != nodeID {
			newChildren = append(newChildren, child)
		}
	}
	currentParent.Children = newChildren

	// Add to new parent
	node.ParentID = newParentID
	newParent.Children = append(newParent.Children, node)

	return SaveMaterials(tree)
}

// UpdateNode updates a node's properties
func UpdateNode(tree *models.MaterialsTree, nodeID string, name, materialType, url, description string) error {
	node := FindNodeByID(tree.Root, nodeID)
	if node == nil {
		return errors.New("nó não encontrado")
	}

	if name != "" {
		node.Name = name
	}
	if node.Type == "material" {
		if materialType != "" {
			node.MaterialType = materialType
		}
		if url != "" {
			node.URL = url
		}
		if description != "" {
			node.Description = description
		}
	}

	return SaveMaterials(tree)
}
