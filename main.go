package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// public struct names are accessible outside of the file
// lowercase stays isolated
type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
	Favorite int    `json:"favorite"`
}

var books = []book{
	{ID: "1", Title: "Code: The Hidden Language of Computer Hardware and Software", Author: "Charles Petzold", Quantity: 20, Favorite: 1},
	{ID: "2", Title: "Cracking the Coding Interview: 189 Programming Questions & Solutions", Author: "Gayle Laakmann McDowell", Quantity: 4, Favorite: 1},
	{ID: "3", Title: "Code Complete", Author: "Steve McConnell", Quantity: 15, Favorite: 0},
	{ID: "4", Title: "Deep Learning", Author: "Ian Goodfellow, Yoshua Bengio and Aaron Courville", Quantity: 27, Favorite: 1},
	{ID: "5", Title: "Grokking Algorithms", Author: "Aditya Y Bhargava", Quantity: 3, Favorite: 0},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id query param"})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "book not available"})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context) {

	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id query param"})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "book not available"})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)

}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	c.IndentedJSON(http.StatusCreated, book)
}

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")

}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	// add the newly created book to the book obj
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)

}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", createBook)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnBook)
	router.Run("localhost:8080")
}
