package routes

import (
	"app/pkg/cryproher"
	"encoding/json"
	"net/http"
	"strconv"
)

func Decrypt(w http.ResponseWriter, r *http.Request) {
	var (
		value   string
		reqJSON []byte
	)

	value = r.URL.Query().Get("value")

	w.Header().Set("Content-Type", "application/json")

	if len(value) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		reqJSON, _ = json.Marshal(map[string]string{
			"method":     r.Method,
			"path":       r.URL.Path,
			"error_code": strconv.Itoa(http.StatusBadRequest),
			"name":       "null params",
		})
		w.Write(reqJSON)
		return
	}

	w.WriteHeader(http.StatusOK)
	reqJSON, _ = json.Marshal(map[string]string{
		"decrypt": cryproher.DecryptLetters(value),
	})
	w.Write(reqJSON)
}

func Encrypt(w http.ResponseWriter, r *http.Request) {
	var (
		value   string
		reqJSON []byte
	)

	value = r.URL.Query().Get("value")

	w.Header().Set("Content-Type", "application/json")

	if len(value) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		reqJSON, _ = json.Marshal(map[string]string{
			"method":     r.Method,
			"path":       r.URL.Path,
			"error_code": strconv.Itoa(http.StatusBadRequest),
			"name":       "null params",
		})
		w.Write(reqJSON)
		return
	}

	w.WriteHeader(http.StatusOK)
	reqJSON, _ = json.Marshal(map[string]string{
		"encrypt": cryproher.EncryptLetter(value),
	})

	w.Write(reqJSON)
}
