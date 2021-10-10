package handler

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	log.Debug("[health] bar")
	w.Write([]byte("{foo:bar}"))
}
