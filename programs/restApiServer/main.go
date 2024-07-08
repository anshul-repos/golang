package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func main() {

	// initializing new router
	router := mux.NewRouter()

	// route incoming requests to path
	router.HandleFunc("/items", getItems).Methods("GET")
	router.HandleFunc("/items/{id}", getItem).Methods("GET")
	router.HandleFunc("/item", addItem).Methods("POST")
	router.HandleFunc("/items/{id}/{name}", updateItems).Methods("PUT")
	router.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")

	// Start http server on 7777 port
	log.Fatal(http.ListenAndServe("localhost:8089", router))

}

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var Items []Item

// gets ALl Items available
func getItems(writer http.ResponseWriter, request *http.Request) {
	json.NewEncoder(writer).Encode(Items)
}

// gets item with requested item id and returns Item. If not found, returns error
func getItem(writer http.ResponseWriter, request *http.Request) {

	reqParams := mux.Vars(request)

	id, err := strconv.Atoi(reqParams["id"])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	for _, item := range Items {
		if item.ID == id {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
	http.NotFound(writer, request)
}

// add item
func addItem(writer http.ResponseWriter, request *http.Request) {

	var item Item

	err := json.NewDecoder(request.Body).Decode(&item)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	item.ID = len(Items) + 1
	Items = append(Items, item)

	json.NewEncoder(writer).Encode(item)
}

func updateItems(writer http.ResponseWriter, request *http.Request) {

	requestParams := mux.Vars(request)

	id, err := strconv.Atoi(requestParams["id"])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	//name := requestParams["name"]

	var requestItemBody Item

	err = json.NewDecoder(request.Body).Decode(&requestItemBody)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	for index, it := range Items {
		if it.ID == id {
			Items[index] = requestItemBody
			json.NewEncoder(writer).Encode(requestItemBody)
			return
		}
	}

	http.NotFound(writer, request)
}

func deleteItem(writer http.ResponseWriter, request *http.Request) {

	reqestParams := mux.Vars(request)
	id, err := strconv.Atoi(reqestParams["id"])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	for i, it := range Items {
		if it.ID == id {
			Items = append(Items[:i], Items[i+1:]...)
			writer.WriteHeader(http.StatusAccepted)
			return
		}
	}
	http.NotFound(writer, request)
}
