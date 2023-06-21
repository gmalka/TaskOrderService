package model

type User struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	Info     UserInfo `json:"info"`
}

type UserAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserInfo struct {
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	Surname   string  `json:"surname"`
	Group     string  `json:"group"`
	Balance   float64 `json:"balance,omitempty"`
}

type UserWithRole struct {
	User User
	Role string
}

type UserInfoWithRole struct {
	Info UserInfo
	Role string
}
