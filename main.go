package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	// "errrors"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost time ", Author: "MArcel Proust", Quantity: 2},
	{ID: "2", Title: "The great GAtsby ", Author: "F.Scoot Fitzuuur", Quantity: 5},
	{ID: "3", Title: "Titanic ", Author: "Jon byden", Quantity: 12},
}

func getbooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func BookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func creatBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return

	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)

}

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing  id query params"})

	}
	book, err := getBookByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return

	}

	if book.Quantity <= 0 {

		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not found"})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing  id query params"})

	}
	book, err := getBookByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return

	}
	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)

}
func getBookByID(id string) (*book, error) {

	for i, b := range books {

		if b.ID == id {
			return &books[i], nil

		}
	}

	return nil, errors.New("book not found")
}

func main() {

	router := gin.Default()
	router.GET("/books", getbooks)
	router.GET("/books/:id", BookById)
	router.POST("/books", creatBook)
	router.PATCH("checkout", checkoutBook)
	router.PATCH("return", returnBook)
	router.Run("localhost:8080")

}
