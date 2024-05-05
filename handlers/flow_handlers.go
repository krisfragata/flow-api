package main

import (
	"net/http"
)

func NotFoundError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "404 not found", http.StatusNotFound)
}