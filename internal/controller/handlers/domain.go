package handlers

type ResponseData struct {
	Success int         `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
