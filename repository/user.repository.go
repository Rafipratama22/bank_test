package repository

import (
	"github.com/Rafipratama22/bank_test/dto"
	"github.com/Rafipratama22/bank_test/entity"
	"github.com/Rafipratama22/bank_test/middleware"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	Register(user entity.UserEntity) (entity.UserEntity, dto.ErrResponse)
	Login(email string, password string) (string, dto.ErrResponse)
	Logout(id uuid.UUID) dto.ErrResponse
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (c *userRepository) Register(user entity.UserEntity) (entity.UserEntity, dto.ErrResponse) {
	// return repo.db.Create(user).Error
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, dto.ErrResponse{Message: "Failed to Generate Password"}
	}
	user.Password = string(hashedPassword)
	result := c.db.Create(&user).Error
	if result != nil {
		return user, dto.ErrResponse{Message: "Failed to Create User"}
	}
	return user, dto.ErrResponse{Message: ""}
}

func (c *userRepository) Login(email string, password string) (string, dto.ErrResponse) {
	var user entity.UserEntity
	var middleware middleware.AuthMiddleware = middleware.NewAuthMiddleware(c.db)
	var userToken string
	err := c.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return "", dto.ErrResponse{Message: "Email or Password is wrong"}
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", dto.ErrResponse{Message: "Email or Password is wrong"}
	}
	userToken, err = middleware.CreateToken(user.ID)
	if err != nil {
		return "", dto.ErrResponse{Message: "Failed to Create Token"}
	}
	result := c.db.Model(&user).Update("is_active", true)
	if result.Error != nil {
		return "", dto.ErrResponse{Message: "Failed to Update User"}
	}
	return userToken, dto.ErrResponse{}
}

func (c *userRepository) Logout(id uuid.UUID) dto.ErrResponse {
	var user entity.UserEntity
	c.db.Where("id = ?", id).First(&user)
	err := c.db.Model(&user).Update("is_active", false).Error
	if err != nil {
		return dto.ErrResponse{Message: "Failed to Update User"}
	}
	return dto.ErrResponse{Message: ""}
}
