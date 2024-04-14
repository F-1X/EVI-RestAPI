package model

import "time"

type Ad struct {
	ID          string    `json:"-" firestore:"-" db:"id"`
	Name        string    `json:"name" firestore:"name" db:"name"`
	Description string    `json:"description,omitempty" firestore:"description" db:"description"`
	Price       float32   `json:"price" firestore:"price" db:"price"`
	CreatedAt   time.Time `json:"-" firestore:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"-" firestore:"updated_at" db:"updated_at"`
}
