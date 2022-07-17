package service

import (
	"encoding/json"
	"net/http"

	"github.com/nurhidaylma/golang-simple-login/config"
	"github.com/nurhidaylma/golang-simple-login/repository"
	"github.com/nurhidaylma/golang-simple-login/repository/postgres"
	"golang.org/x/crypto/bcrypt"
)

// Login is service for user to login
func Login(w http.ResponseWriter, r *http.Request) {
	connection := config.GetDatabase()
	defer config.Closedatabase(connection)

	// Get the request parameter
	var authdetails repository.Authentication
	err := json.NewDecoder(r.Body).Decode(&authdetails)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "server error", 500)
		return
	}

	// Check the existing user
	existingUser, err := postgres.NewUserPostgres().FindByEmail(authdetails.Email)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "server error", 500)

		return
	}
	if existingUser == nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "user is not registered", 400)

		return
	}
	// Compare the password from request parameter with the actual one
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(authdetails.Password))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "invalid credentials", 400)

		return
	}

	// Generate JWT
	validToken, err := config.GenerateJWT(existingUser.Email, existingUser.Role)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "failed to generate token", 500)

		return
	}

	var token repository.Token
	token.Email = existingUser.Email
	token.Role = existingUser.Role
	token.TokenInString = validToken
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}
