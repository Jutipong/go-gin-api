package model

type ReponseModel struct {
	Status  bool        `json:"status" default:false`
	Code    int         `json:"code" default:500`
	Message string      `json:"message"`
	Datas   interface{} `json:"datas`
}