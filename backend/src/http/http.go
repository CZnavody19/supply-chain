package http

import (
	"net/http"

	"github.com/CZnavody19/supply-chain/src/db"
)

type HttpHandler struct {
	store *db.DatabaseStore
}

func NewHttpHandler(store *db.DatabaseStore) *HttpHandler {
	return &HttpHandler{
		store: store,
	}
}

func (hh *HttpHandler) Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("meow :3"))
}
