package bookstore

import (
	"errors"
	"fmt"
)

type Book struct {
	Title  string
	Author string
	Copies int
}

func Buy(b Book) (Book, error) {
	if b.Copies < 1 {
		return Book{}, errors.New("no copies left")
	}

	b.Copies--

	return b, nil
}

func GetAllBooks(catalog []Book) []Book {
	return catalog
}

func GetBook(catalog map[int]Book, ID int) (Book, error) {
	book, ok := catalog[ID]

	if !ok {
		return Book{}, fmt.Errorf("ID %d doesn't exist", ID)
	}

	return book, nil
}
