package services

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"moses/models"
	"moses/repository"
)

// =========================
// REGISTER
// =========================

func Register(name, email, password string) error {

	// check existing user
	existing, _ := repository.GetUserByEmail(email)
	if existing != nil {
		return errors.New("email already registered")
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Role:     "ADMIN", // default role (bisa nanti RBAC)
	}

	return repository.CreateUser(&user)
}

// =========================
// LOGIN
// =========================

func Login(email, password string) (*models.User, error) {

	user, err := repository.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}
