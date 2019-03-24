package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}

var users []User

func listUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func main() {
	router := mux.NewRouter()
	users = append(users, User{Username: "afhdfgbdf", Email: "adfbdfbdfbd@a.a"})
	users = append(users, User{Username: "bdfbdfbdf", Email: "bfbdfbdfbdfb@b.b"})
	users = append(users, User{Username: "cbdfbdfbfbdfb"})
	router.HandleFunc("/", listUsers)
	log.Fatal(http.ListenAndServe(":12345", router))
}
