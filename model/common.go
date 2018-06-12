package model


const DeletedAt = "0000-01-01 00:00:00"

type QueryParams struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}