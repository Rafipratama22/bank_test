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

// Register
// @Summary Register yang dilakukan oleh user baik dia sebagai consumer maupun merchant untuk pengembalian data nya akan mengembalikan user yang sudah terdaftar dengan password yang sudah di hash
// @Description Register yang dilakukan oleh user baik dia sebagai consumer maupun merchant untuk pengembalian data nya akan mengembalikan user yang sudah terdaftar dengan password yang sudah di hash
// @Tags Form
// @Accept  */*
// @Produce  json
// @Param data body entity.UserEntity true "User"
// @Success 201 {object} entity.UserEntity
// @Failure 400 {object} dto.ErrResponse
// @Router /user/register [post]
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

// Login
// @Summary Login yang dilakukan setelah user melakukan register, jika sudah berhasil maka pengguna akan mendapatkan token yang akan digunakan untuk mengakses API
// @Description Login yang dilakukan setelah user melakukan register, jika sudah berhasil maka pengguna akan mendapatkan token yang akan digunakan untuk mengakses API
// @Tags Form
// @Accept  */*
// @Produce  json
// @Param data body dto.LoginDto true "User"
// @Success 201 {object} dto.LoginDtoResponse
// @Failure 400 {object} dto.ErrResponse
// @Router /user/login [post]
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

// Logout
// @Summary Logout dapat dilakukan oleh user yang sudah login, jika sudah berhasil
// @Description Logout dapat dilakukan oleh user yang sudah login, jika sudah berhasil
// @Tags Form
// @Accept  */*
// @Produce  json
// @Security ApiKeyAuth
// @Param data body dto.LoginDto true "User"
// @Success 201 {object} dto.LogoutDtoResponse
// @Failure 400 {object} dto.ErrResponse
// @Router /user/logout [post]
func (c *userController) Logout(ctx *gin.Context) {
	var errResponse dto.ErrResponse
	var logoutResponse dto.LogoutDtoResponse
	userId := ctx.MustGet("user_id")
	userID := uuid.MustParse(userId.(string))
	errResponse = c.userRepo.Logout(userID)
	if errResponse.Message != "" {
		errResponse.Message = "Failed to Logout User"
		ctx.JSON(http.StatusBadRequest, errResponse)
	} else {
		logoutResponse.Message = "Successfully Logout User"
		ctx.JSON(http.StatusOK, logoutResponse)
	}
}
