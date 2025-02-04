package main

import (
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

var books = []Book{
	{Id: "1", Title: "The Go Programming Language", Author: "Alan A. A. Donovan", Price: 34.99},
}
var mu sync.Mutex

func main() {
	route := gin.Default()
	route.GET("/", homeHandler)
	route.GET("/books", getBooks)
	route.POST("/books", createBook)
	route.GET("/books/:id", getBookById)
	route.PUT("/books/:id", updateBook)
	route.DELETE("/books/:id", deleteBook)

	// server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	route.Run(":" + port)
}

func homeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Book APIs",
	})
}
func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var newbook Book
	if err := c.BindJSON(&newbook); err != nil {
		return
	}

	mu.Lock()
	books = append(books, newbook)
	mu.Unlock()

	c.JSON(http.StatusCreated, newbook)
}

func getBookById(c *gin.Context) {
	id := c.Param("id")

	for _, book := range books {
		if book.Id == id {
			c.JSON(http.StatusOK, book)
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Book Not found!"})
}

func updateBook(c *gin.Context) {
	id := c.Param("id")
	var updatebook Book
	if err := c.BindJSON(&updatebook); err != nil {
		return
	}

	mu.Lock()
	for i, book := range books {
		if book.Id == id {
			books[i] = updatebook
			c.JSON(http.StatusOK, updatebook)
			mu.Unlock()
			return
		}
	}
	mu.Unlock()
	c.JSON(http.StatusNotFound, gin.H{"message": "Book Not Found!"})
}

func deleteBook(c *gin.Context) {
	id := c.Param("id")
	mu.Lock()
	for i, book := range books {
		if book.Id == id {
			books = append(books[:1], books[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
			mu.Unlock()
			return
		}
	}
	mu.Unlock()
	c.JSON(http.StatusNotFound, gin.H{"message": "Book Not Found!"})
}
