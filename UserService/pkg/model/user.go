package model

type User struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	Info     UserInfo `json:"info"`
}

type UserWithoutBalance struct {
	Username string                 `json:"username"`
	Password string                 `json:"password"`
	Info     UserInfoWithoutBalance `json:"info"`
}

type UserForUpdate struct {
	Username string            `json:"username"`
	Password string            `json:"password"`
	Info     UserInfoForUpdate `json:"info"`
}

type UserAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserInfoForUpdate struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Surname   string `json:"surname"`
	Group     string `json:"group"`
}

type UserWithRole struct {
	User User
	Role string
}

type UserInfoWithRole struct {
	Info UserInfo
	Role string
}

type UserInfo struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Surname   string `json:"surname"`
	Group     string `json:"group"`
	Balance   int    `json:"balance"`
}

type UserInfoWithoutBalance struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Surname   string `json:"surname"`
	Group     string `json:"group"`
}
