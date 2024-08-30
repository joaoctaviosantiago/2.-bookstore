package bookstore

import (
	"errors"
	"fmt"
)

type Book struct {
	ID              int
	Title           string
	Author          string
	Copies          int
	PriceCents      int
	DiscountPercent int
}

type Catalog map[int]Book

func Buy(b Book) (Book, error) {
	if b.Copies < 1 {
		return Book{}, errors.New("no copies left")
	}

	b.Copies--

	return b, nil
}

func (c Catalog) GetAllBooks() []Book {
	result := []Book{}

	for _, b := range c {
		result = append(result, b)
	}

	return result
}

func (c Catalog) GetBook(ID int) (Book, error) {
	book, ok := c[ID]

	if !ok {
		return Book{}, fmt.Errorf("ID %d doesn't exist", ID)
	}

	return book, nil
}

func (b Book) NetPriceCents() int {
	discountValue := (b.PriceCents * b.DiscountPercent) / 100

	return b.PriceCents - discountValue
}

func (b *Book) SetPriceCents(newPrice int) error {
	if newPrice < 0 {
		return fmt.Errorf("negative price %d not allowed", newPrice)
	}

	b.PriceCents = newPrice
	return nil
}
