CREATE TABLE user (
    id INTEGER PRIMARY KEY,
    user_name string NOT NULL,
    password string NOT NULL,
    email string NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    
);