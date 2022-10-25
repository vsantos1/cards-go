package repository

import (
	"log"
	"net/http"

	"github.com/vsantos1/bank-go/database"
	"github.com/vsantos1/bank-go/models"
)

type UserRepository struct {
	User *models.User `json:"user"`
	Err error `json:"error"`
	Status int `json:"status_code"`
	Users []models.User `json:"users"`
}

func GetById (id int) UserRepository {
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	var user models.User
	err = db.Get(&user,"SELECT * FROM user WHERE id = ?", id)
	
	if err != nil {
		return UserRepository{
			User: nil,
			Err: err,
			Status: http.StatusNotFound,
		}
	}

	return UserRepository{
		User: &user,
		Err: err,
		Status: http.StatusOK,
	}
}

func Create(u *models.User) UserRepository {

	db,_ := database.Connect()
	
	_,err := db.Exec("INSERT INTO user (user_name, pass_word, email, balance) VALUES (?, ?, ?, ?)", u.Username,u.Password, u.Email, u.Balance)
	if err != nil {
		return UserRepository{
			Err: err,
			Status: http.StatusBadRequest,
			User: nil,
		}
	}

	defer db.Close()
	
	return UserRepository{
		User: u,
		Err: err,
		Status: http.StatusCreated,
	}
}

func UpdateBalance(id int,u *models.User) UserRepository {

	db,_ := database.Connect()
	
	_,err := db.Exec("UPDATE user SET user_name = ?, pass_word = ?, email = ?, balance = ? WHERE id = ?", u.Username,u.Password, u.Email, u.Balance, id)
	if err != nil {
		return UserRepository {
			Err: err,
			Status: http.StatusBadRequest,
			User: nil,
		}
	}

	defer db.Close()
	
	return UserRepository{
		User: u,
		Err: err,
		Status: http.StatusOK,}
}

func Delete(id int ) UserRepository{
	db,_ := database.Connect()
	
	_,err := db.Exec("DELETE FROM user WHERE id = ?", id)
	if err != nil {
		return UserRepository{
			Err: err,
			Status: http.StatusBadRequest,
			User: nil,
		}
	}

	defer db.Close()
	
	return UserRepository{
		Err: err,
		Status: http.StatusNoContent,
	}
}

func GetAll() UserRepository{
	db,_ := database.Connect()
	
	var users []models.User
	err := db.Select(&users,"SELECT * FROM user")
	if err != nil {
		return UserRepository{
			Err: err,
			Status: http.StatusBadRequest,
			User: nil,
		}
	}

	defer db.Close()
	
	return UserRepository{
		Users: users,
		Status: http.StatusOK,
	}
}