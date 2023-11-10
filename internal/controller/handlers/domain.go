package handlers

type ResponseData struct {
	Success int    `json:"success"`
	Message string `json:"message"`
	Id      string `json:"id,omitempty"`
}
