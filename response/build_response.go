package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Meta    MetaInfo    `json:"meta"`
}

type MetaInfo struct {
	HttpStatus int `json:"http_status"`
	Offset     int `json:"offset,omitempty"`
	Limit      int `json:"limit,omitempty"`
	Total      int `json:"total,omitempty"`
}

func Write(w http.ResponseWriter, response Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Meta.HttpStatus)
	json.NewEncoder(w).Encode(&response)
}

func OKWithMeta(w http.ResponseWriter, data interface{}, message string, meta MetaInfo) {
	meta.HttpStatus = 200
	response := Response{
		Data:    data,
		Message: message,
		Meta:    meta,
	}
	Write(w, response)
}

func OK(w http.ResponseWriter, data interface{}, message string) {
	meta := MetaInfo{}
	OKWithMeta(w, data, message, meta)
}

func Error(w http.ResponseWriter, err error, httpStatus int) {
	meta := MetaInfo{
		HttpStatus: httpStatus,
	}
	response := Response{
		Message: err.Error(),
		Meta:    meta,
	}
	Write(w, response)
}
