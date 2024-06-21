package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func sendData(rw http.ResponseWriter, data interface{}, status int) {
	rw.Header().Set("content-type", "application/json")
	rw.WriteHeader(status)
	ouput, _ := json.Marshal(data)
	fmt.Fprint(rw, string(ouput))
}

func sendError(rw http.ResponseWriter, err error, status int) {
	rw.WriteHeader(status)
	ouput, _ := json.Marshal(err)
	fmt.Fprint(rw, "Error: "+string(ouput))
}
