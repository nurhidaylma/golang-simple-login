package service

import (
	"encoding/json"
	"net/http"

	"github.com/nurhidaylma/golang-simple-login/config"
	"github.com/nurhidaylma/golang-simple-login/repository"
	"github.com/nurhidaylma/golang-simple-login/repository/postgres"
	"golang.org/x/crypto/bcrypt"
)

// Register is the service for user to register an account
func Register(w http.ResponseWriter, r *http.Request) {
	connection := config.GetDatabase()
	defer config.Closedatabase(connection)

	var user repository.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "server error", 500)

		return
	}

	// Check for the existing user
	existingUser, err := postgres.NewUserPostgres().FindByEmail(user.Email)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "server error", 500)
		return
	}
	if existingUser != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "user already exists", 400)
		return
	}

	// Hash the password
	bcryptHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "server error", 500)
		return
	}

	// Insert user data into the database
	insertedUser, err := postgres.NewUserPostgres().Insert(&repository.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: string(bcryptHash),
		Role:     user.Role,
	})
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "server error", 500)
		return
	}
	if insertedUser == nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "invalid request", 500)
		return
	}

}
