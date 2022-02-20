package controller

import (
	"net/http"

	"github.com/Rafipratama22/bank_test/dto"
	"github.com/Rafipratama22/bank_test/entity"
	"github.com/Rafipratama22/bank_test/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MerchantController interface{
	CreateMerchant(ctx *gin.Context)
	GetMerchantAll(ctx *gin.Context)
	DetailMerchant(ctx *gin.Context)
}

type merchantController struct {
	merchantRepo repository.MerchantRepository
}

func NewMerchantController(merchantRepo repository.MerchantRepository) MerchantController {
	return &merchantController{
		merchantRepo: merchantRepo,
	}
}

func (m *merchantController) CreateMerchant(ctx *gin.Context) {
	var errResponse dto.ErrResponse
	var merchant entity.MerchantEntity
	userId := ctx.MustGet("user_id")
	userID := uuid.MustParse(userId.(string))
	err := ctx.ShouldBindJSON(&merchant)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse)
	}
	merchant.UserID = userID
	result, errResponse := m.merchantRepo.CreateMerchant(merchant)
	if errResponse.Message != "" {
		ctx.JSON(http.StatusBadRequest, errResponse)
	} else {
		ctx.JSON(http.StatusCreated, result)
	}
}

func (m *merchantController) GetMerchantAll(ctx *gin.Context) {
	var errResponse dto.ErrResponse
	userId := ctx.MustGet("user_id")
	userID := uuid.MustParse(userId.(string))
	result, errResponse := m.merchantRepo.GetMerchantAll(userID)
	if errResponse.Message != "" {
		ctx.JSON(http.StatusBadRequest, errResponse)
	} else {
		ctx.JSON(http.StatusOK, result)
	}
}

func (m *merchantController) DetailMerchant(ctx *gin.Context) {
	var errResponse dto.ErrResponse
	merchantId := ctx.Param("id")
	merchantID := uuid.MustParse(merchantId)
	result, errResponse := m.merchantRepo.DetailMerchant(merchantID)
	if errResponse.Message != "" {
		ctx.JSON(http.StatusBadRequest, errResponse)
	} else {
		ctx.JSON(http.StatusOK, result)
	}
}