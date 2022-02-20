package dto

type UpdateRekening struct {
	Balance int `json:"balance"`
	Pin string `json:"pin"`
}

type TransferRekening struct {
	NoRekening string `json:"no_rekening"`
	Balance int `json:"balance"`
	Pin string `json:"pin"`
	ID int `json:"id"`
}