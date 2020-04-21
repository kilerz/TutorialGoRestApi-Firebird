package _struct

type ResponseData struct {
	Status int64 `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}
