package bookstore

import "errors"

type Book struct {
	Title  string
	Author string
	Copies int
}

func Buy(b Book) (Book, error) {
	if b.Copies < 1 {
		return Book{}, errors.New("mo copies left")
	}

	b.Copies--

	return b, nil
}
