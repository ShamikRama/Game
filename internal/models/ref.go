package models

type RefInput struct {
	Id int `json:"referral_code" binding:"required"`
}
