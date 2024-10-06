package main

import (
	"fmt"
	"strconv"
)

type Stringer interface {
	String() string
}

type Count int

type Book struct {
	Title  string
	Author string
}

func (book Book) string() string {
	return book.Title + " by " + book.Author
}

func (count Count) string() string {
	return strconv.Itoa(int(count))
}

func main() {
	fmt.Println("---> Interface <---")
	i, err := strconv.Atoi("-11")

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	book := Book{"The Art of Computer Programming", "Donald Knuth"}

	fmt.Println("Book: ", book)

	fmt.Println("Converted integer: ", i)
}
