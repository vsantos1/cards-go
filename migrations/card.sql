
CREATE TABLE card (
  	id INTEGER PRIMARY KEY,
    user_id INTEGER,
    card_owner string NOT NULL,
    card_number int NOT NULL,
    card_type string NOT NULL,
    card_cvv int NOT NULL,
    card_expiry DATE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user(id)
  )
