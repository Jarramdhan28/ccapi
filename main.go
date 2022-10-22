package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type student struct {
	ID          string
	Name        string
	Member      string
	Profession  string
	SosialMedia string
}

var data = []student{
	{"CC0001", "Zam-Zam", "Pasif", "Full Timer Ruang Guru", "@zamzam"},
	{"CC0002", "Ipan Badruzaman", "Aktif", "Lead SE", "@ipan"},
	{"CC0003", "Muhamad Fajar", "Aktif", "Co Lead SE", "@fajar"},
	{"CC0004", "Alfiani FK", "Aktif", "VueJs", "@alfi"},
	{"CC0005", "Moch Fikri", "Pasif", "-", "@fikri"},
}

func users(w http.ResponseWriter, r *http.Request) {
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

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var id = r.FormValue("id")
		var result []byte
		var err error

		for _, each := range data {
			if each.ID == id {
				result, err = json.Marshal(each)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Write(result)
				return
			}
		}

		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/users", users)
	http.HandleFunc("/user", user)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
