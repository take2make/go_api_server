package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type ResponsePostData struct {
  Detail int64
}

var usercounter int = 0

var activeusers map[int]string

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
}

func ResHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	if activeusers[id] != "" {
		res := fmt.Sprint(`{"status": `, id, `}`)
		w.Write([]byte(res))
	} else {
		w.Write([]byte("Not in work"))
	}
}

func post(w http.ResponseWriter, r *http.Request) {
	data := ResponsePostData{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("detail = %d", data.Detail)

	usercounter += 1
	activeusers[usercounter] = "In work"

	resp := fmt.Sprintf(`{"session": %d}`, usercounter)
	w.Write([]byte(resp))
	if usercounter == 1 {
		for {

		}
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}

func main() {
	activeusers = make(map[int]string)
	r := mux.NewRouter()
	r.HandleFunc("/", get).Methods(http.MethodGet)
	r.HandleFunc("/", post).Methods(http.MethodPost)
	r.HandleFunc("/", notFound)
	r.HandleFunc("/{id:[0-9]+}", ResHandler).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", r))
}
