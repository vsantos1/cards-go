package models

type User struct {
	ID        int     `json:"id" db:"id" `
	Username  string  `json:"user_name" db:"user_name"`
	Password  string  `json:"-" db:"pass_word"`
	Email     string  `json:"email" db:"email"`
	Balance   float64 `json:"balance" db:"balance"`
	CreatedAt string  `json:"created_at" db:"created_at"`
	UpdatedAt string  `json:"updated_at" db:"updated_at"`
}

type Users []User