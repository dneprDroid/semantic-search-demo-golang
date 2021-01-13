package responseutil 

import (
	"net/http"
	"encoding/json"
)

func WriteJson(w http.ResponseWriter, object interface{}) error {
	resp, err := json.Marshal(object)
	if err != nil {
		return err 
	}
	w.Write(resp)
	return nil 
}

func Error(w http.ResponseWriter, err error) {
	ErrorStr(w, err.Error())
}

func ErrorStr(w http.ResponseWriter, errStr string) {
	w.WriteHeader(http.StatusInternalServerError)
	errMsg := map[string]string {
		"error": errStr,
	}
	WriteJson(w, errMsg)
}