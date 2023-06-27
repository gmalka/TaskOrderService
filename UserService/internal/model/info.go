package model

type AuthInfo struct {
	Access  string `json:"Access"`
	Refresh string `json:"Refresh"`
}

type ResponseMessage struct {
	Message string `json:"message"`
}
