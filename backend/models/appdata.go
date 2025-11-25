package models

// Reminder represents a reminder/task item
type Reminder struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Date      string `json:"date"`
	Time      string `json:"time"`
	Priority  string `json:"priority"`
	Completed bool   `json:"completed"`
}

// Note represents a study note
type Note struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Subject string `json:"subject"`
	Date    string `json:"date"`
}

// Event represents a calendar event
type Event struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Date        string `json:"date"`
	Time        string `json:"time"`
	Description string `json:"description"`
}

// Material represents a study material
type Material struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Folder      string `json:"folder"`
	URL         string `json:"url"`
	DateAdded   string `json:"dateAdded"`
}

// AppData represents all application data for a user
type AppData struct {
	Notes     []Note         `json:"notes"`
	Reminders []Reminder     `json:"reminders"`
	StudyLog  map[string]int `json:"studyLog"`
	Subjects  []string       `json:"subjects"`
	Events    []Event        `json:"events"`
	Materials []Material     `json:"materials"`
	Folders   []string       `json:"folders"`
}
