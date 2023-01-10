package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
	Audition string `json:"audition"`
}

var books = []book{
	{Id: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2, Audition: "10th"},
	{Id: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5, Audition: "11th"},
	{Id: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6, Audition: "12th"},
	{Id: "4", Title: "Allama Iqbal", Author: "Sheraz khan", Quantity: 7, Audition: "18th"},
	{Id: "5", Title: "Lost Men", Author: "Robert GreezeMen", Quantity: 4, Audition: "1st"},
}

func returnBook(c *gin.Context) {

	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id"})
		return
	}

	book, err := GetBookById(id)

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)

}

func checkout(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id"})
		return
	}

	book, err := GetBookById(id)

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not available"})
		return
	}
	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)

}

func GetBookId(c *gin.Context) {

	id := c.Param("id")
	book, err := GetBookById(id)
	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)

}
func GetBookById(id string) (*book, error) {
	for k, v := range books {
		if v.Id == id {
			return &books[k], nil
		}
	}

	return nil, errors.New("book not found")

}
func AddBook(c *gin.Context) {
	var newBook book
	err := c.BindJSON(&newBook)
	if err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func GetBook(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func main() {

	router := gin.Default()
	router.GET("/books", GetBook)
	router.PATCH("/checkout", checkout)
	router.PATCH("/return", returnBook)
	router.GET("books/:id", GetBookId)
	router.POST("/books", AddBook)
	router.Run("localhost:8080")

}
