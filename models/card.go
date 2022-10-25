package models

import (
	"time"
)

type Card struct {
	ID         int     `json:"id"`
	UserId    int     `json:"user_id"`
	CardOwner  string  `json:"card_owner"`
	CardLimit  float64 `json:"card_limit"`
	CardNumber int     `json:"card_number"`
	CardType  string  `json:"card_type"`
	CardCvv    int     `json:"card_cvv"`
	CardExpiry time.Time `json:"card_expiry"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`


}