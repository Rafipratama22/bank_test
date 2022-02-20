package controller

import (
	"net/http"
	"strconv"

	"github.com/Rafipratama22/bank_test/dto"
	"github.com/Rafipratama22/bank_test/entity"
	"github.com/Rafipratama22/bank_test/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RekeningController interface{
	CreateRekening(ctx *gin.Context)
	AllRekening(ctx *gin.Context)
	GetRekeningByID(ctx *gin.Context)
	UpdateRekening(ctx *gin.Context)
	TransferRekening(ctx *gin.Context)
}

type rekeningController struct {
	rekeningRepo repository.RekeningRepository
}

func NewRekeningController(rekeningRepo repository.RekeningRepository) RekeningController {
	return &rekeningController{
		rekeningRepo: rekeningRepo,
	}
}

func (s *rekeningController) CreateRekening(ctx *gin.Context) {
	var saving entity.RekeningEntity
	var errResponse dto.ErrResponse
	userId := ctx.MustGet("user_id")
	userID := uuid.MustParse(userId.(string)) 
	err := ctx.ShouldBindJSON(&saving)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	saving.UserID = userID
	errResponse = s.rekeningRepo.CreateRekening(saving)
	if errResponse.Message != "" {
		ctx.JSON(http.StatusBadRequest, errResponse)
	} else {
		ctx.JSON(http.StatusCreated, gin.H{
			"message": "Rekening Created",
		})
	}
}

func (s *rekeningController) AllRekening(ctx *gin.Context) {
	var errResponse dto.ErrResponse
	var rekening []entity.RekeningEntity
	userId := ctx.MustGet("user_id")
	userID := uuid.MustParse(userId.(string)) 
	rekening, errResponse = s.rekeningRepo.AllRekening(userID)
	if errResponse.Message != "" {
		ctx.JSON(http.StatusBadRequest, errResponse)
	} else {
		ctx.JSON(http.StatusOK, rekening)
	}
}

func (s *rekeningController) GetRekeningByID(ctx *gin.Context) {
	var errResponse dto.ErrResponse
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	rekening, errResponse := s.rekeningRepo.GetRekeningByID(id)
	if errResponse.Message != "" {
		ctx.JSON(http.StatusBadRequest, errResponse)
	} else {
		ctx.JSON(http.StatusOK, rekening)
	}
}

func (s *rekeningController) UpdateRekening(ctx *gin.Context) {
	var errResponse dto.ErrResponse
	var saving dto.UpdateRekening
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	err = ctx.ShouldBindJSON(&saving)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	errResponse = s.rekeningRepo.UpdateRekening(id, saving.Balance, saving.Pin)
	if errResponse.Message != "" {
		ctx.JSON(http.StatusBadRequest, errResponse)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Rekening Updated",
		})
	}
}

func (s *rekeningController) TransferRekening(ctx *gin.Context) {
	var errResponse dto.ErrResponse
	var saving dto.TransferRekening
	err := ctx.ShouldBindJSON(&saving)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	// saving.NoRekening = (uuid.MustParse(saving.NoRekening))
	noRekening := uuid.MustParse(saving.NoRekening)
	errResponse = s.rekeningRepo.TransferRekening(noRekening, saving.Balance, saving.Pin, saving.ID)
	if errResponse.Message != "" {
		if errResponse.Message == "Failed because commit error" {
			ctx.JSON(http.StatusInternalServerError, errResponse)
		} else {
			ctx.JSON(http.StatusBadRequest, errResponse)
		}
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Rekening Transferred",
		})
	}
}