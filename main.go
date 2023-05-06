package main

import (
	"fmt"
	"strconv"

	//"log"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	FullNme  string
	UserName string
	Email    string
}

// the Post is a struct  that  represent  a single  post, which is the instance of the user

type Post struct {
	Title  string
	Body   string
	Author User
}

var data []Post = []Post{}  //global variable data, which is a slice of Post. This will be used to store the posts created through the API

func main() {

	router := mux.NewRouter()

	//router.HandleFunc("/test" ,test)
	//router.HandleFunc("/add/{item}",addItems).Methods("GET", "DELETE")
	router.HandleFunc("/posts", getItem).Methods("GET")
	router.HandleFunc("/posts", addItems).Methods("POST")
	router.HandleFunc("/posts/{id}", getPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", updateItem).Methods("PUT")
	router.HandleFunc("/posts/{id}", patchItem).Methods("PATCH")
	router.HandleFunc("/posts/{id}", deleteItem).Methods("DELETE")
	http.ListenAndServe(":8080", router)

}

func getPosts(w http.ResponseWriter, r *http.Request) {   //extracts the ID parameter from the URL using mux.Vars(),if the ID parameter cannot be converted to an integer, or if the ID is out of range, it returns an appropriate error message,if the ID is valid, it returns the corresponding post using json.NewEncoder()
	w.Header().Set("Content-Type", "Application/json")

	var idParam string = mux.Vars(r)["id"]

	id, err := strconv.Atoi(idParam)
	if err != nil {

		w.WriteHeader(400)
		w.Write([]byte("ID could  not be  converted  to  integer"))
		return
	}

	if id >= len(data) {
		w.WriteHeader(404)
		w.Write([]byte("No data found with  specified ID"))
		return
	}

	post := data[id]

	json.NewEncoder(w).Encode(post)
}

func addItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	var newPost Post

	json.NewDecoder(r.Body).Decode(&newPost)

	data = append(data, newPost)

	json.NewEncoder(w).Encode(data)

}

func getItem(w http.ResponseWriter, r *http.Request) {   //getItem() function handles GET requests to the "/posts" endpoint. It sets the response header to JSON and encodes the data slice using json.NewEncoder()

	w.Header().Set("Content-Type", "Application/json")

	fmt.Println("Your details")

	json.NewEncoder(w).Encode(data)

}

func updateItem(w http.ResponseWriter, r *http.Request) {    //extracting the ID parameter and checking its validity, then decodes the request body into a new Post instance, replaces the corresponding element in the data slice with the new instance, and encodes the updated instance back to the response using json.NewEncoder()
	w.Header().Set("Content-Type", "Application/json")

	var idParam string = mux.Vars(r)["id"]

	id, err := strconv.Atoi(idParam)

	if err != nil {

		w.WriteHeader(400)
		w.Write([]byte("ID could not converted to Integer"))
		return
	}

	//error checking

	if id >= len(data) {

		w.WriteHeader(404)
		w.Write([]byte("No data founded with  specified ID"))
		return

	}

	var updatedItem Post

	//updateItem := Post

	json.NewDecoder(r.Body).Decode(&updatedItem)

	data[id] = updatedItem

	json.NewEncoder(w).Encode(updatedItem)

}

func patchItem(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "Application/json")

	var idParam string = mux.Vars(r)["id"]

	id, err := strconv.Atoi(idParam)

	if err != nil {

		w.WriteHeader(400)
		w.Write([]byte("ID could not converted to Integer"))
		return
	}

	//error checking

	if id >= len(data) {

		w.WriteHeader(404)
		w.Write([]byte("No data founded with  specified ID"))
		return

	}

	// get the  current  value

	patchdata := data[id]

	json.NewDecoder(r.Body).Decode(&patchdata)

	data[id] = patchdata

	json.NewEncoder(w).Encode(patchdata)

}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	var idParam string = mux.Vars(r)["id"]

	id, err := strconv.Atoi(idParam)

	if err != nil {

		w.WriteHeader(400)
		w.Write([]byte("ID could not converted to Integer"))
		return
	}

	//error checking

	if id >= len(data) {

		w.WriteHeader(404)
		w.Write([]byte("No data founded with  specified ID"))
		return

	}

	data = append(data[:id], data[id+1:]...)

	w.WriteHeader(200)
}
