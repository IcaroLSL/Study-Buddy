package storage

import (
	"studybuddy/models"
	"time"
)

// FindSubjectByID busca recursivamente uma matéria pelo ID na árvore
func FindSubjectByID(subjects []models.Subject, id int64) *models.Subject {
	for i := range subjects {
		if subjects[i].ID == id {
			return &subjects[i]
		}
		if found := FindSubjectByID(subjects[i].Children, id); found != nil {
			return found
		}
	}
	return nil
}

// AddSubject adiciona uma matéria na posição correta da árvore
// Se parentID for nil, adiciona na raiz
// Retorna true se a matéria foi adicionada com sucesso
func AddSubject(subjects *[]models.Subject, subject models.Subject, parentID *int64) bool {
	if parentID == nil {
		*subjects = append(*subjects, subject)
		return true
	}

	return addSubjectToParent(subjects, subject, *parentID)
}

// addSubjectToParent busca o pai e adiciona a matéria como filho
func addSubjectToParent(subjects *[]models.Subject, subject models.Subject, parentID int64) bool {
	for i := range *subjects {
		if (*subjects)[i].ID == parentID {
			(*subjects)[i].Children = append((*subjects)[i].Children, subject)
			return true
		}
		if addSubjectToParent(&(*subjects)[i].Children, subject, parentID) {
			return true
		}
	}
	return false
}

// RemoveSubject remove uma matéria pelo ID (e todos os seus filhos)
// Retorna true se a matéria foi removida com sucesso
func RemoveSubject(subjects *[]models.Subject, id int64) bool {
	for i := range *subjects {
		if (*subjects)[i].ID == id {
			*subjects = append((*subjects)[:i], (*subjects)[i+1:]...)
			return true
		}
		if RemoveSubject(&(*subjects)[i].Children, id) {
			return true
		}
	}
	return false
}

// MoveSubject move uma matéria para outro pai
// Se newParentID for nil, move para a raiz
// Retorna true se a matéria foi movida com sucesso
func MoveSubject(subjects *[]models.Subject, subjectID int64, newParentID *int64) bool {
	// Primeiro, encontrar e remover a matéria atual
	subject := findAndRemoveSubject(subjects, subjectID)
	if subject == nil {
		return false
	}

	// Atualizar o parentID da matéria
	subject.ParentID = newParentID

	// Adicionar no novo local
	return AddSubject(subjects, *subject, newParentID)
}

// findAndRemoveSubject encontra, remove e retorna uma matéria
func findAndRemoveSubject(subjects *[]models.Subject, id int64) *models.Subject {
	for i := range *subjects {
		if (*subjects)[i].ID == id {
			subject := new(models.Subject)
			*subject = (*subjects)[i]
			*subjects = append((*subjects)[:i], (*subjects)[i+1:]...)
			return subject
		}
		if found := findAndRemoveSubject(&(*subjects)[i].Children, id); found != nil {
			return found
		}
	}
	return nil
}

// FlattenSubjects retorna uma lista plana de todas as matérias
func FlattenSubjects(subjects []models.Subject) []models.Subject {
	var result []models.Subject
	for _, subject := range subjects {
		// Adiciona a matéria atual (sem os filhos no item plano)
		flatSubject := models.Subject{
			ID:          subject.ID,
			Name:        subject.Name,
			Description: subject.Description,
			ParentID:    subject.ParentID,
			Color:       subject.Color,
			Icon:        subject.Icon,
			CreatedAt:   subject.CreatedAt,
		}
		result = append(result, flatSubject)
		// Adiciona os filhos recursivamente
		result = append(result, FlattenSubjects(subject.Children)...)
	}
	return result
}

// BuildSubjectTree reconstrói a árvore a partir de uma lista plana
func BuildSubjectTree(flatSubjects []models.Subject) []models.Subject {
	// Mapa de ID para Subject
	subjectMap := make(map[int64]*models.Subject)
	for i := range flatSubjects {
		flatSubjects[i].Children = []models.Subject{}
		subjectMap[flatSubjects[i].ID] = &flatSubjects[i]
	}

	var roots []models.Subject
	for i := range flatSubjects {
		if flatSubjects[i].ParentID == nil {
			roots = append(roots, flatSubjects[i])
		} else {
			parent, exists := subjectMap[*flatSubjects[i].ParentID]
			if exists {
				parent.Children = append(parent.Children, flatSubjects[i])
			}
		}
	}

	return roots
}

// UpdateSubjectInTree atualiza uma matéria existente na árvore
// Retorna true se a matéria foi atualizada com sucesso
func UpdateSubjectInTree(subjects *[]models.Subject, updated models.Subject) bool {
	for i := range *subjects {
		if (*subjects)[i].ID == updated.ID {
			// Preserva os filhos atuais
			updated.Children = (*subjects)[i].Children
			(*subjects)[i] = updated
			return true
		}
		if UpdateSubjectInTree(&(*subjects)[i].Children, updated) {
			return true
		}
	}
	return false
}

// GetNextSubjectID retorna o próximo ID disponível para uma nova matéria
// Usa timestamp como ID
func GetNextSubjectID() int64 {
	return time.Now().UnixMilli()
}
