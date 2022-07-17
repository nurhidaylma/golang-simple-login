package repository

type Token struct {
	Role          string `json:"role"`
	Email         string `json:"email"`
	TokenInString string `json:"token"`
}
