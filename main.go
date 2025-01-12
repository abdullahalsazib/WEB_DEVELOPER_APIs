package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Define a model for your data
type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var db *gorm.DB
var err error

// Initialize the database connection
func init() {
	// Open a MySQL connection using GORM
	dsn := "root:1234@tcp(localhost:3306)/my_database?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Could not connect to the database: ", err)
	}
	fmt.Println("Successfully connected to MySQL database!")

	// Migrate the User model to the database
	db.AutoMigrate(&User{})
}

func main() {
	// Initialize the router
	r := mux.NewRouter()

	// Define routes and associate them with handler functions
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/users/{eid}", getUserByID).Methods("GET")
	r.HandleFunc("/users/first", getTheFirstUser).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")

	// Start the server on port 8080
	log.Fatal(http.ListenAndServe(":8080", r))
}

// Handle the endpoint to retrieve all users
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

// Handle the endpoint to retrieve a user by ID
func getUserByID(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the URL params
	w.Header().Set("Content-Type", "application/json")

	// Parse the ID from the URL (eid)
	eid := mux.Vars(r)["eid"]
	id, err := strconv.Atoi(eid) // Convert string to int
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var user User
	// Query the database for the user by ID
	if err := db.First(&user, id).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Return the user as JSON
	json.NewEncoder(w).Encode(user)
}

// Handle the endpoint to retrieve the first user
func getTheFirstUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User

	// Retrieve the first user from the database
	if err := db.First(&user).Error; err != nil {
		http.Error(w, "No users found", http.StatusNotFound)
		return
	}

	// Return the first user as JSON
	json.NewEncoder(w).Encode(user)
}

// Handle the endpoint to create a new user
func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	// Decode the request body into the user struct
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Create the user in the database
	if err := db.Create(&user).Error; err != nil {
		http.Error(w, "Could not create user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User created with ID: %d", user.ID)
}
