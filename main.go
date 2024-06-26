package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Myuser struct {
	UID  string
	Name string
}

var data = []Myuser{
	Myuser{"90124", "Jon12"}, Myuser{"94996", "627User"}, Myuser{"12324", "Regsy23"},
}

func OUH(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}
func none1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		fmt.Printf("There's nothing to see here\nthere is no one1")

		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}
func main() {
	http.HandleFunc("/ouh", OUH)
	http.HandleFunc("/nonee", none1)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
