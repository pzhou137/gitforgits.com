package main

import (
	"fmt"
)

type Book struct {
	ID int
	Title string
	Author string
	Description string
	Price float64
}

func getBookDetails(id int) *Book {
	book := &Book{
		ID: id,
		Title: "The Go Programming Language",
		Author: "Alan A. A. Donovan and Brian W. Kernighan",
		Description: "The Go Programming Language is a statically typed, compiled language that supports multiple programming paradigms.",
		Price: 49.99,
	}
	return book
}

func main() {
	book := getBookDetails(1)
	fmt.Printf("Book Details: %+v\n", book)
}
