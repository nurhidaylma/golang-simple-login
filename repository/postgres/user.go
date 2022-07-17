package postgres

import (
	"errors"

	"github.com/google/uuid"
	"github.com/nurhidaylma/golang-simple-login/config"
	"github.com/nurhidaylma/golang-simple-login/repository"
	"gorm.io/gorm"
)

// userPostgres implements the user repository service
type userPostgres struct {
	db *gorm.DB
}

// FindByEmail get user by email
func (s *userPostgres) FindByEmail(email string) (*repository.User, error) {
	db := s.db

	user := repository.User{}
	err := db.First(&user, `"email" = ? AND "deleted_at" IS NULL`, email).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Insert inserts a user
func (s *userPostgres) Insert(singleUser *repository.User) (*repository.User, error) {
	db := s.db

	singleUser.ID, _ = uuid.NewRandom()

	var err error

	if err = db.Create(singleUser).Error; err != nil {
		return nil, err
	}

	return singleUser, nil
}

// NewUserPostgres creates new user repository service
func NewUserPostgres() repository.UserRepository {
	return &userPostgres{
		db: config.GetDatabase(),
	}
}
