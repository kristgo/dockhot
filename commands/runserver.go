// Pushing created docker container to Github
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the gorilla/mux API.\nIt is the best API on Golang.")
	})

	fmt.Println("Server listening!")
	http.ListenAndServe(":80", r)
}
