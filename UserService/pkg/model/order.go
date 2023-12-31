package model

type TaskOrderInfo struct {
	Id     int `db:"id" json:"id,omitempty"`
	Answer int `db:"answer" json:"answer,omitempty"`
	Price  int `db:"price" json:"price,omitempty"`
}

type Task struct {
	Id      int     `db:"id" json:"id,omitempty"`
	Count   int     `db:"count" json:"count"`
	Heights []int64 `db:"heights" json:"heights"`
	Price   int     `db:"price" json:"price"`
	Answer  int     `db:"price" json:"answer"`
}

type TaskWithoutAnswer struct {
	Id      int     `db:"id" json:"id,omitempty"`
	Count   int     `db:"count" json:"count"`
	Heights []int64 `db:"heights" json:"heights"`
	Price   int     `db:"price" json:"price"`
}

type TaskAnswer struct {
	Answer int `json:"answer"`
}

type BalanceChange struct {
	Username string `json:"username"`
	Money    int    `json:"money"`
}
