package models

// User represents a user in the system
type User struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}

// LoginRequest represents the login request body
type LoginRequest struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	RememberMe bool   `json:"rememberMe"`
}

// AuthResponse represents the authentication response
type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
