package server

import (
	"api/database/methods"
	"api/structs"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func StartServer() {

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			return
		}

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
		w.Write(methods.GetAll())
	})

	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			return
		}

		id := strings.TrimPrefix(r.URL.Path, "/users/")

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
		w.Write(methods.Get(id))
	})

	http.HandleFunc("/users/{userId}/edit", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			return
		}

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")

		var user structs.EditUser

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusOK)
		methods.Edit(user.Id, user)
	})

	fmt.Println("Listening to server...")
	http.ListenAndServe(":8080", nil)
}
