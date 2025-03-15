package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
}

var userCache = make(map[int]User)

// https://www.youtube.com/watch?v=eqvDSkuBihs
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("POST /users", createUser)

	fmt.Println("Server is listening at http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}

func handleRoot(
	w http.ResponseWriter,
	r *http.Request,
) {
	// printf places output on the standard output stream stdout.
	// fprintf places output on the named output stream.
	// https://www.sas.upenn.edu/~saul/parasite/man/man3/printf.3.html#:~:text=printf%20places%20output%20on%20the,that%20enough%20storage%20is%20available.
	fmt.Fprintf(w, "Hello World!")
}

func createUser(
	w http.ResponseWriter,
	r *http.Request,
) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	userCache[len(userCache)+1] = user

	w.WriteHeader(http.StatusNoContent)
}
