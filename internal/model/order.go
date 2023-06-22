package model

type Order struct {
	Id      int     `db:"id" json:"id,omitempty"`
	Count   int     `db:"count" json:"count"`
	Heights []int64 `db:"heights" json:"heights"`
	Price   int     `db:"price" json:"price"`
}