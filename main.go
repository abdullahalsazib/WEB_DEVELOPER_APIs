package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type RouteResponse struct {
	Message string `json:"message"`
	ID      string `json:"id,omitempty"`
}

func main() {
	fmt.Println("web-developer-project")

	router := mux.NewRouter()

	router.HandleFunc("/register", register).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/projects", createProject).Methods("POST")
	router.HandleFunc("/projects/{id}", updateProject).Methods("PUT")
	router.HandleFunc("/projects", getProject).Methods("GET")
	router.HandleFunc("/projects/{id}", getProject).Methods("GET")
	router.HandleFunc("/projects/{id}", deleteProject).Methods("DELETE")

	// server on serve
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}

// register
func register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(RouteResponse{Message: "Hello, Message from register!"})
	if err != nil {
		log.Fatal(err)
	}
}

// login
func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(RouteResponse{Message: "Hello, Message from login!"})
	if err != nil {
		log.Fatal(err)
	}
}

// createProject
func createProject(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(RouteResponse{Message: "Hello, Message from createProject!"})
	if err != nil {
		log.Fatal(err)
	}
}

// updateProject
func updateProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(RouteResponse{Message: "Hello, Message from updateProject!", ID: id})
	CheckTheError(err)
}

// getProject
func getProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(RouteResponse{Message: "Hello, Message from getProject!", ID: id})
	if err != nil {
		log.Fatal(err)
	}
}

// deleteProject
func deleteProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(RouteResponse{Message: "Hello, Message from deleteProject!", ID: id})
	if err != nil {
		log.Fatal(err)
	}
}

func CheckTheError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
