package model

type AuthInfo struct {
	Access  string `json:"Refresh"`
	Refresh string `json:"Access"`
}
