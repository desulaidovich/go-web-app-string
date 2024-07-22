package routes

import (
	"app/pkg/cryproher"
	"encoding/json"
	"net/http"
	"strconv"
)

type Route struct {
	*http.ServeMux
	cryproher.Cryproher
}

func New(s *http.ServeMux) *Route {
	var crypto cryproher.Cryproher

	return &Route{
		s, crypto,
	}
}

func (s *Route) Decrypt(w http.ResponseWriter, r *http.Request) {
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
		"decrypt": s.DecryptLetters(value),
	})
	w.Write(reqJSON)
}

func (s *Route) Encrypt(w http.ResponseWriter, r *http.Request) {
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

	letters := s.EncryptLetter(value)
	lettersExp := s.EncryptExpr(letters)

	w.WriteHeader(http.StatusOK)
	reqJSON, _ = json.Marshal(map[string]string{
		"encrypt": lettersExp,
	})
	w.Write(reqJSON)
}
