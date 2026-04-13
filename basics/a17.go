package main

import "fmt"

type Book struct {
	title string
	auth  string
}

func changebook(book *Book) {
	book.auth = "cc"
}

func main() {
	var book Book
	book.title = "aa"
	book.auth = "bb"
	fmt.Println(book)

	changebook(&book)
	fmt.Println(book)
}
