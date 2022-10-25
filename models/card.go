package models

import (
	"time"
)

type Card struct {
	ID         int     `json:"id" db:"id"`
	UserId    int     `json:"user_id" db:"user_id"`
	CardOwner  string  `json:"card_owner" db:"card_owner"`
	CardLimit  float64 `json:"card_limit" db:"card_limit"`
	CardNumber int     `json:"card_number" db:"card_number"`
	CardType  string  `json:"card_type" db:"card_type"`
	CardCvv    int     `json:"card_cvv" db:"card_cvv"`
	CardExpiry time.Time `json:"card_expiry" db:"card_expiry"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`


}

type Cards []Card