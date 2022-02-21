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
	GetRekeningAll(ctx *gin.Context)
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

// Rekening Create
// @Summary user yang sudah login dapat membuka rekening baru dengan memasukan pin saja karena user akan terhubung dengan rekening setelah dibuat
// @Description user yang sudah login dapat membuka rekening baru dengan memasukan pin saja karena user akan terhubung dengan rekening setelah dibuat
// @Tags Rekening
// @Accept  */*
// @Produce  json
// @Param data body entity.RekeningEntity true "Rekening"
// @Security ApiKeyAuth
// @Success 201 {object} dto.RegisterRekening
// @Failure 400 {object} dto.ErrResponse
// @Router /rekening/create [post]
func (s *rekeningController) CreateRekening(ctx *gin.Context) {
	var saving entity.RekeningEntity
	var errResponse dto.ErrResponse
	var result dto.RegisterRekening
	userId := ctx.MustGet("user_id")
	userID := uuid.MustParse(userId.(string)) 
	err := ctx.ShouldBindJSON(&saving)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	saving.UserID = userID
	result, errResponse = s.rekeningRepo.CreateRekening(saving)
	if errResponse.Message != "" {
		ctx.JSON(http.StatusBadRequest, errResponse)
	} else {
		ctx.JSON(http.StatusCreated, result)
	}
}

// Rekening Get All
// @Summary user yang sudah login dapat mengakses rekening yang telah dibuat, sistem akan menyocokan rekening yang telah dibuat oleh user
// @Description user yang sudah login dapat mengakses rekening yang telah dibuat, sistem akan menyocokan rekening yang telah dibuat oleh user
// @Tags Rekening
// @Accept  */*
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} []entity.RekeningEntity
// @Failure 400 {object} dto.ErrResponse
// @Router /rekening/list [get]
func (s *rekeningController) GetRekeningAll(ctx *gin.Context) {
	var errResponse dto.ErrResponse
	var rekening []entity.RekeningEntity
	userId := ctx.MustGet("user_id")
	userID := uuid.MustParse(userId.(string)) 
	rekening, errResponse = s.rekeningRepo.GetRekeningAll(userID)
	if errResponse.Message != "" {
		ctx.JSON(http.StatusBadRequest, errResponse)
	} else {
		ctx.JSON(http.StatusOK, rekening)
	}
}

// Rekening Get By ID
// @Summary user yang sudah login dapat mengakses rekening yang telah dibuat dengan memberi id rekening
// @Description user yang sudah login dapat mengakses rekening yang telah dibuat dengan memberi id rekening
// @Tags Rekening
// @Accept  */*
// @Produce  json
// @Security ApiKeyAuth
// @Param id path int true "id"
// @Success 200 {object} entity.RekeningEntity
// @Failure 400 {object} dto.ErrResponse
// @Router /rekening/detail/:id [get]
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

// Update Rekening
// @Summary user yang sudah login dapat mengupdate balance yang ke rekening yang sudah ada dengan memasukkan pin dan sejumlah uang yang ingin dimasukkan
// @Description user yang sudah login dapat mengupdate balance yang ke rekening yang sudah ada dengan memasukkan pin dan sejumlah uang yang ingin dimasukkan
// @Tags Rekening
// @Accept  */*
// @Produce  json
// @Security ApiKeyAuth
// @Param id path int true "id"
// @Param data body dto.UpdateRekening true "Rekening"
// @Success 200 {object} dto.UpdateRekeningResponse
// @Failure 400 {object} dto.ErrResponse
// @Router /rekening/update/:id [put]
func (s *rekeningController) UpdateRekening(ctx *gin.Context) {
	var errResponse dto.ErrResponse
	var saving dto.UpdateRekening
	var response dto.UpdateRekeningResponse
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
		response.Message = "Rekening Updated"
		ctx.JSON(http.StatusOK, response)
	}
}

// Transfer Rekening
// @Summary user yang sudah login baik customer maupun merchant dapat melakukan transfer uamg ke rekening user yang telah terdaftar, note untuk id di dto merupakan id rekening yang kita punya
// @Description user yang sudah login baik customer maupun merchant dapat melakukan transfer uamg ke rekening user yang telah terdaftar, note untuk id di dto merupakan id rekening yang kita punya
// @Tags Rekening
// @Accept  */*
// @Produce  json
// @Security ApiKeyAuth
// @Param data body dto.TransferRekening true "Rekening"
// @Success 200 {object} dto.TransferRekeningResponse
// @Failure 400 {object} dto.ErrResponse
// @Router /rekening/transfer [post]
func (s *rekeningController) TransferRekening(ctx *gin.Context) {
	var errResponse dto.ErrResponse
	var saving dto.TransferRekening
	err := ctx.ShouldBindJSON(&saving)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	userId := ctx.MustGet("user_id")
	userID := uuid.MustParse(userId.(string)) 
	noRekening := uuid.MustParse(saving.NoRekening)
	errResponse = s.rekeningRepo.TransferRekening(noRekening, saving.Balance, saving.Pin, saving.ID, userID)
	if errResponse.Message != "" {
		if errResponse.Message == "Failed because commit error" {
			ctx.JSON(http.StatusInternalServerError, errResponse)
		} else {
			ctx.JSON(http.StatusBadRequest, errResponse)
		}
	} else {
		ctx.JSON(http.StatusOK, dto.TransferRekeningResponse{
			Message: "Transfer Success",
			Balance: saving.Balance,
			NoRekening: saving.NoRekening,
		})
	}
}

