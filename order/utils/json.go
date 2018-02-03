package utils

import (
	"encoding/json"
	"net/http"
)

// JsonResponse function for Marshal and format response using Json
func JsonResponse(res http.ResponseWriter, resp interface{}, httpCode int) {
	msg, _ := json.Marshal(resp)
	res.Header().Set("Content-Type", "application-json; charset=utf-8")
	res.WriteHeader(httpCode)
	res.Write(msg)
}
