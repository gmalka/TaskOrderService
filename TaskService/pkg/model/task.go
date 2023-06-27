package model

type Task struct {
	Id      int     `db:"id" json:"id,omitempty"`
	Count   int     `db:"quantity" json:"quantity"`
	Heights []int64 `db:"heights" json:"heights"`
	Price   int     `db:"price" json:"price"`
	Answer  int     `db:"answer" json:"answer"`
}

type UsersPurchase struct {
	Username string `json:"username"`
	OrderId  int    `json:"orderId"`
}

type TaskNewPrice struct {
	Id    int `json:"orderId"`
	Price int `json:"price"`
}

type TaskWithoutAnswer struct {
	Id      int     `db:"id" json:"id,omitempty"`
	Count   int     `db:"quantity" json:"quantity"`
	Heights []int64 `db:"heights" json:"heights"`
	Price   int     `db:"price" json:"price"`
}