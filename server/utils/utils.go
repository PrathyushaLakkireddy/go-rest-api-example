package utils

import (
	"encoding/json"
	"net/http"

	"github.com/globalsign/mgo"

	"github.com/PrathyushaLakkireddy/go-rest-api-example/server/common"
)

func ReturnSuccessReponse(statusCode int, msg string, data interface{}, res http.ResponseWriter) {
	res.Header().Set("Content-Type", "application/json")

	res.WriteHeader(statusCode)

	_ = json.NewEncoder(res).Encode(common.Response{
		Success: true,
		Message: msg,
		Data:    data,
	})
	return
}

func ReturnErrorResponse(statusCode int, msg string, notfoundMsg string, err error, data interface{},
	res http.ResponseWriter) {

	res.Header().Set("Content-Type", "application/json")
	if err == nil {
		res.WriteHeader(statusCode)
		_ = json.NewEncoder(res).Encode(common.Response{
			Success: false,
			Message: msg,
			Error:   nil,
		})
		return
	}

	if err == mgo.ErrNotFound {
		res.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(res).Encode(common.Response{
			Success: false,
			Message: msg,
			Error:   err,
		})
		return
	} else {
		res.WriteHeader(statusCode)
		_ = json.NewEncoder(res).Encode(common.Response{
			Success: false,
			Message: msg,
			Error:   err,
		})
		return
	}
}
