package repository

import (
	"log"
	"net/http"

	"github.com/vsantos1/bank-go/database"
	"github.com/vsantos1/bank-go/models"
	"github.com/vsantos1/bank-go/utils"
)

type CardRepository struct {
	Card *models.Card `json:"card"`
	Err error `json:"error"`
	Status int `json:"status_code"`
	Cards []models.Card `json:"cards"`
}

func GetCardById (id int) CardRepository {
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	var card models.Card
	err = db.Get(&card,"SELECT * FROM card WHERE id = ?", id)
	
	if err != nil {
		return CardRepository{
			Card: nil,
			Err: err,
			Status: http.StatusNotFound,
		}
	}

	return CardRepository{
		Card: &card,
		Err: err,
		Status: http.StatusOK,
	}
}

func GetCardFK(user_id int ) CardRepository {
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	var card models.Card
	err = db.Get(&card,"SELECT * FROM card WHERE user_id = ?", user_id)

	if err != nil {
		return CardRepository{
			Card: nil,
			Err: err,
			Status: http.StatusNotFound,
		}
	}

	return CardRepository{
		Card: &card,
		Err: err,
		Status: http.StatusOK,
	}
}

func CreateCard(id int,c *models.Card) CardRepository {

	db,_ := database.Connect()


	
	_,err := db.Exec("INSERT INTO card (user_id, card_owner,card_limit, card_number, card_type,card_cvv,card_expiry) VALUES (?,?, ?, ?, ?,?,?)", id,c.CardOwner,c.CardLimit,utils.Dec(12),c.CardType,utils.Dec(3),utils.ExpiryCardDate())
	if err != nil {
		return CardRepository{
			Err: err,
			Status: http.StatusBadRequest,
			Card: nil,
		}
	}

	defer db.Close()
	
	return CardRepository{
		Card: c,
		Err: err,
		Status: http.StatusCreated,
	}
}

