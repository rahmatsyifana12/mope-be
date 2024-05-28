package dtos

type Response struct {
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}