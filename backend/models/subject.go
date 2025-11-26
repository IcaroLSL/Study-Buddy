package models

// Subject representa uma matéria de estudo em estrutura de árvore
type Subject struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ParentID    *int64    `json:"parentId,omitempty"` // nil para matérias raiz
	Children    []Subject `json:"children,omitempty"` // sub-matérias
	Color       string    `json:"color,omitempty"`    // cor para UI
	Icon        string    `json:"icon,omitempty"`     // ícone opcional
	CreatedAt   string    `json:"createdAt"`
}
