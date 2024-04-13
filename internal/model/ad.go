package model

import "time"

type Ad struct {
	Name        string    `json:"name" firestore:"name"`
	Description string    `json:"description,omitempty" firestore:"description"`
	Price       float32   `json:"price" firestore:"price"`
	CreatedAt   time.Time `json:"-" firestore:"created_at"`
	UpdatedAt   time.Time `json:"-" firestore:"updated_at"`
}

type AdsResponse struct {
	Ads []*Ad `json:"ads"`
}
