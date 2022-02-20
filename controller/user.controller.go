package controller

import (
	"net/http"

	"github.com/Rafipratama22/bank_test/dto"
	"github.com/Rafipratama22/bank_test/entity"
	"github.com/Rafipratama22/bank_test/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
}

type userController struct {
	userRepo repository.UserRepository
}

func NewUserController(userRepo repository.UserRepository) UserController {
	return &userController{
		userRepo: userRepo,
	}
}

func (c *userController) Register(ctx *gin.Context) {
	var user entity.UserEntity
	var errResponse dto.ErrResponse
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		panic(err)
	}
	usered, errResponse := c.userRepo.Register(user)
	if errResponse.Message != "" {
		ctx.JSON(http.StatusBadRequest, errResponse)
	} else {
		ctx.JSON(http.StatusCreated, usered)
	}
}

func (c *userController) Login(ctx *gin.Context) {
	var loginDto dto.LoginDto
	var errResponse dto.ErrResponse
	err := ctx.ShouldBindJSON(&loginDto)
	if err != nil {
		errResponse.Message = "Failed to Bind"
		ctx.JSON(http.StatusInternalServerError, errResponse)
	}
	token, errResponse := c.userRepo.Login(loginDto.Email, loginDto.Password)
	if token == "" {
		if errResponse.Message == "Failed to Create Token" {
			ctx.JSON(http.StatusInternalServerError, errResponse)
		}
		ctx.JSON(http.StatusBadRequest, errResponse)
	} else {
		ctx.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func (c *userController) Logout(ctx *gin.Context) {
	var errResponse dto.ErrResponse
	userId := ctx.MustGet("user_id")
	userID := uuid.MustParse(userId.(string))
	errResponse = c.userRepo.Logout(userID)
	if errResponse.Message != "" {
		errResponse.Message = "Failed to Logout User"
		ctx.JSON(http.StatusBadRequest, errResponse)
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Logout User Success"})
	}
}
