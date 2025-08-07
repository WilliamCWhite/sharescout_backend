package handlers

import (
	"net/http"
)

// Simply for ensuring that the backend is receiving requests
func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Backend is working"))
}
