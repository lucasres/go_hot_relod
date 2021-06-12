package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HelloResponse struct {
	To string
}

func main() {
	http.HandleFunc("/", IndexFunc)

	fmt.Println("start server at 8000")
	http.ListenAndServe(":8000", nil)
}

func IndexFunc(w http.ResponseWriter, r *http.Request) {
	response := HelloResponse{
		To: "🚽😠",
	}

	json.NewEncoder(w).Encode(response)
}
