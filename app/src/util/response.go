package util 

import (
	"net/http"
	"encoding/json"
)

func RespWriteJson(w http.ResponseWriter, object interface{}) error {
	response, err := json.Marshal(object)
	if err != nil {
		return err 
	}
	w.Write(response)
	return nil 
}

func RespError(w http.ResponseWriter, err error) {
	RespErrorStr(w, err.Error())
}

func RespErrorStr(w http.ResponseWriter, errStr string) {
	w.WriteHeader(http.StatusInternalServerError)
	errMsg := map[string]string {
		"error": errStr,
	}
	RespWriteJson(w, errMsg)
}