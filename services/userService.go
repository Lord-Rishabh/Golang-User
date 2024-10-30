package services

import (
	"errors"
	"user_service/database"
	"user_service/models"

	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    if err != nil {
        fmt.Println("Password comparison failed:", err)
    }
    return err == nil
}


func Signup(user models.User) (models.User, error) {
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return user, err
	}
	user.Password = hashedPassword
	result := database.DB.Create(&user)
	return user, result.Error
}

func Login(email string, password string) (*models.User, error) {
	var user models.User
	result := database.DB.Where("email = ?", email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}
	if !CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid password")
	}
	return &user, nil
}

func GetUser(id string) (*models.User, error) {
	var user models.User
	result := database.DB.Where("id = ?", id).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := database.DB.Find(&users)
	return users, result.Error
}