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
	UpdateMerchant(ctx *gin.Context)
	DeleteMerchant(ctx *gin.Context)
}

type merchantController struct {
	merchantRepo repository.MerchantRepository
}

func NewMerchantController(merchantRepo repository.MerchantRepository) MerchantController {
	return &merchantController{
		merchantRepo: merchantRepo,
	}
}

// Create Merchant
// @Summary user yang sudah login dapat membuat akun merchant untuk kebutuhan pribadi, dengan memasukan no_rekening, address, dan lain lain
// @Description user yang sudah login dapat membuat akun merchant untuk kebutuhan pribadi, dengan memasukan no_rekening, address, dan lain lain
// @Tags Merchant
// @Accept  */*
// @Produce  json
// @Security ApiKeyAuth
// @Param data body entity.MerchantEntity true "Merchant"
// @Success 201 {object} entity.MerchantEntity
// @Failure 400 {object} dto.ErrResponse
// @Router /merchant/create [post]
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

// Get Merchant All
// @Summary user yang sudah login dapat melihat semua merchant yang ada di bank dan dapat melihat detail merchant seperti user dan rekening
// @Description user yang sudah login dapat melihat semua merchant yang ada di bank dan dapat melihat detail merchant seperti user dan rekening
// @Tags Merchant
// @Accept  */*
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} []dto.MerchantListResponse
// @Failure 400 {object} dto.ErrResponse
// @Router /merchant/list [get]
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

// Get Merchant Detail
// @Summary user yang sudah login dapat melihat semua merchant yang ada di bank dan dapat melihat detail merchant seperti user dan rekening
// @Description user yang sudah login dapat melihat semua merchant yang ada di bank dan dapat melihat detail merchant seperti user dan rekening
// @Tags Merchant
// @Accept  */*
// @Produce  json
// @Security ApiKeyAuth
// @Param id path string true "Merchant ID"
// @Success 200 {object} dto.MerchantListResponse
// @Failure 400 {object} dto.ErrResponse
// @Router /merchant/detail/:id [get]
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


// Update Merchant
// @Summary user yang sudah login dan memiliki merchant yang dapat mengupdate merchant tersebut
// @Description user yang sudah login dan memiliki merchant yang dapat mengupdate merchant tersebut
// @Tags Merchant
// @Accept  */*
// @Produce  json
// @Security ApiKeyAuth
// @Param id path string true "Merchant ID"
// @Param data body entity.MerchantEntity true "Merchant"
// @Success 200 {object} entity.MerchantEntity
// @Failure 400 {object} dto.ErrResponse
// @Failure 404 {object} dto.ErrResponse
// @Router /merchant/update/:id [put]
func (m *merchantController) UpdateMerchant(ctx *gin.Context) {
	var errResponse dto.ErrResponse
	var merchant entity.MerchantEntity
	var response dto.UpdateMerchantResponse
	id := ctx.Param("id")
	merchantID := uuid.MustParse(id)
	userId := ctx.MustGet("user_id")
	userID := uuid.MustParse(userId.(string))
	err := ctx.ShouldBindJSON(&merchant)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse)
	}
	merchant.UserID = userID
	_, errResponse = m.merchantRepo.UpdateMerchant(merchantID, merchant)
	if errResponse.Message != "" {
		if errResponse.Message == "Failed to Get Merchant" {
			ctx.JSON(http.StatusNotFound, errResponse)
		} else {
			ctx.JSON(http.StatusBadRequest, errResponse)
		}
	} else {
		response.Message = "Merchant Updated"
		ctx.JSON(http.StatusOK, response)
	}
}

// Delete Merchant
// @Summary user yang sudah login dapat menghapus merchant yang ada di bank
// @Description user yang sudah login dapat menghapus merchant yang ada di bank
// @Tags Merchant
// @Accept  */*
// @Produce  json
// @Security ApiKeyAuth
// @Param id path string true "Merchant ID"
// @Success 200 {object} dto.MerchantListResponse
// @Failure 400 {object} dto.ErrResponse
// @Failure 404 {object} dto.ErrResponse
// @Router /merchant/delete/:id [delete]
func (m *merchantController) DeleteMerchant(ctx *gin.Context) {
	var errResponse dto.ErrResponse
	var response dto.DeleteMerchantResponse
	id := ctx.Param("id")
	merchantID := uuid.MustParse(id)
	userId := ctx.MustGet("user_id")
	userID := uuid.MustParse(userId.(string))
	errResponse = m.merchantRepo.DeleteMerchant(merchantID, userID)
	if errResponse.Message != "" {
		if errResponse.Message == "Failed to Get Merchant" {
			ctx.JSON(http.StatusNotFound, errResponse)
		} else if errResponse.Message == "Not Allowed" {
			ctx.JSON(http.StatusMethodNotAllowed, errResponse)
		} else {
			ctx.JSON(http.StatusBadRequest, errResponse)
		}
	} else {
		response.Message = "Success Delete Merchant"
		ctx.JSON(http.StatusOK, response)
	}
}