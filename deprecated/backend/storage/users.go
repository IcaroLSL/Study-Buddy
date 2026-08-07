package storage

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"studybuddy/models"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var (
	usersFile  = "storage/users.json"
	users      = make(map[string]models.User)
	usersMutex sync.Mutex
	jwtSecret  []byte
)

func init() {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Println("WARNING: JWT_SECRET not set, using default secret for development")
		secret = "default-dev-secret-change-in-production"
	}
	jwtSecret = []byte(secret)
}

// GetJWTSecret returns the JWT secret
func GetJWTSecret() []byte {
	return jwtSecret
}

// LoadUsers loads users from the JSON file
func LoadUsers() {
	usersMutex.Lock()
	defer usersMutex.Unlock()

	bytes, err := os.ReadFile(usersFile)
	if err != nil {
		log.Printf("INFO: Could not read users file: %v (this is normal on first run)", err)
		return
	}
	if err := json.Unmarshal(bytes, &users); err != nil {
		log.Printf("WARNING: Could not parse users file: %v", err)
	}
}

// SaveUsers saves users to the JSON file
func SaveUsers() {
	bytes, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		log.Printf("ERROR: Could not serialize users: %v", err)
		return
	}
	if err := os.WriteFile(usersFile, bytes, 0644); err != nil {
		log.Printf("ERROR: Could not write users file: %v", err)
	}
}

// GetUser retrieves a user by email
func GetUser(email string) (models.User, bool) {
	usersMutex.Lock()
	defer usersMutex.Unlock()
	user, exists := users[email]
	return user, exists
}

// UserExists checks if a user exists by email
func UserExists(email string) bool {
	usersMutex.Lock()
	defer usersMutex.Unlock()
	_, exists := users[email]
	return exists
}

// CreateUser creates a new user
func CreateUser(email, password, name string) (models.User, error) {
	usersMutex.Lock()
	defer usersMutex.Unlock()

	hashedPassword, err := HashPassword(password)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		ID:        time.Now().Unix(),
		Email:     email,
		Password:  hashedPassword,
		Name:      name,
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	users[email] = user
	SaveUsers()
	return user, nil
}

// HashPassword hashes a password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash compares a password with a hash
func CheckPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

// GenerateJWT generates a JWT token for a user
func GenerateJWT(user models.User, rememberMe bool) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	if rememberMe {
		expirationTime = time.Now().Add(7 * 24 * time.Hour)
	}

	claims := &jwt.RegisteredClaims{
		Subject:   fmt.Sprintf("%d", user.ID),
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
