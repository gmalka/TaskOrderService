package model

type Order struct {
	Count int `db:"count" json:"count"`
	Heights []int `db:"heights" json:"heights"`
	Price float64 `db:"price" json:"price"`
}