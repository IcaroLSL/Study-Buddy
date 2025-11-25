package storage

import (
	"encoding/json"
	"os"
	"studybuddy/models"
	"sync"
)

var (
	dataFile  = "storage/data.json"
	dataMutex sync.Mutex
)

// LoadData loads application data from the JSON file
func LoadData() (models.AppData, error) {
	dataMutex.Lock()
	defer dataMutex.Unlock()

	bytes, err := os.ReadFile(dataFile)
	if err != nil {
		return models.AppData{}, nil // Return empty data if file doesn't exist
	}

	var data models.AppData
	if err := json.Unmarshal(bytes, &data); err != nil {
		return models.AppData{}, err
	}
	return data, nil
}

// SaveData saves application data to the JSON file
func SaveData(data models.AppData) error {
	dataMutex.Lock()
	defer dataMutex.Unlock()

	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, bytes, 0644)
}

// DeleteEvent deletes an event by ID and returns the updated data
func DeleteEvent(eventID int64) (models.AppData, error) {
	dataMutex.Lock()
	defer dataMutex.Unlock()

	bytes, err := os.ReadFile(dataFile)
	if err != nil {
		return models.AppData{}, err
	}

	var data models.AppData
	if err := json.Unmarshal(bytes, &data); err != nil {
		return models.AppData{}, err
	}

	newEvents := make([]models.Event, 0)
	for _, event := range data.Events {
		if event.ID != eventID {
			newEvents = append(newEvents, event)
		}
	}

	data.Events = newEvents

	updatedBytes, _ := json.MarshalIndent(data, "", "  ")
	if err := os.WriteFile(dataFile, updatedBytes, 0644); err != nil {
		return models.AppData{}, err
	}

	return data, nil
}
