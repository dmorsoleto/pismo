package handlers

import (
	"encoding/json"
	"net/http"
)

func ParserHttpResponse(w http.ResponseWriter, success int, httpStatus int, message string, data ...interface{}) {
	response := ResponseData{
		Success: success,
		Message: message,
	}

	if data != nil {
		response.Data = data[0]
	}

	responseJson, _ := json.Marshal(response)
	w.WriteHeader(httpStatus)
	w.Write(responseJson)
}
