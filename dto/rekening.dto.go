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

type RegisterRekening struct {
	NoRekening string `json:"no_rekening"`
	Balance int `json:"balance"`
}

type UpdateRekeningResponse struct {
	Message string `json:"message"`
}

type TransferRekeningResponse struct {
	Message string `json:"message"`
	Balance int `json:"balance"`
	NoRekening string `json:"no_rekening"`
}