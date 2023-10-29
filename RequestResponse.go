package main

import (
	"encoding/json"
	"net/http"
)

type ApiResponse struct {
	Header *ApiResponseHeader
	Data   *ApiResponseData
}

type ApiResponseHeader struct {
	HttpStatus string `json:"httpStatus"`
}

type ApiResponseData struct {
	Message   string   `json:"message,omitempty"`
	ActorList []*Actor `json:"ActorList,omitempty"`
}

func ServiceJSONResponse(w http.ResponseWriter, httpStatus string, msg string, responseData interface{}) {
	ApiResponseObject := &ApiResponse{
		Header: &ApiResponseHeader{},
		Data:   &ApiResponseData{},
	}
	ApiResponseObject.Header.HttpStatus = httpStatus
	// Add responseData in json out object
	if msg == "" {
		switch rd := responseData.(type) {
		case []*Actor:
			ApiResponseObject.Data.ActorList = rd
		// Add error response
		case error:
			ApiResponseObject.Data.Message = rd.Error()
		default:
			ApiResponseObject.Data.Message = UnknownResponseType
		}

	} else {
		ApiResponseObject.Data.Message = msg
	}
	// json output for API
	json.NewEncoder(w).Encode(ApiResponseObject)
}
