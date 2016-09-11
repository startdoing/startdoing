package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Home page
func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

// Router creates the http router
func Router() *mux.Router {
	rt := mux.NewRouter()
	rt.HandleFunc("/", Home)
	return rt
}

func main() {
	rt := Router()
	http.ListenAndServe(":8080", rt)
}
