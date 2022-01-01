package server

import (
	"encoding/json"
	"net/http"
)

func JSONResponse(status int, val any, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	out, _ := json.Marshal(val)
	_, _ = w.Write(out)
}

func ErrorResponse(status int, err error, w http.ResponseWriter) {
	JSONResponse(status, struct {
		Err string `json:"err"`
	}{
		Err: err.Error(),
	}, w)
}
