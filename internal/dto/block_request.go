package dto

type BlockRequest struct {
	Height int `json:"height" validate:"required"`
	Weight int `json:"weight" validate:"required"`
}
