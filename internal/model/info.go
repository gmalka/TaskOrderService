package model

type AuthInfo struct {
	Access  string `json:"Refresh"`
	Refresh string `json:"Access"`
}

type ResponseMessage struct {
	Message string `json:"message"`
}
