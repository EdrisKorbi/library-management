package models

import (
	"errors"
	"fmt"
)

type Book struct {
	ID        int
	Title     string
	Author    string
	ISBN      string
	Available bool
}

func (b Book) DisplayInfo() string {
	return fmt.Sprintf(
		"ID: %d | Title: %s | Author: %s | ISBN: %s | Available: %t",
		b.ID, b.Title, b.Author, b.ISBN, b.Available,
	)
}

func (b *Book) Borrow() error {
	if !b.Available {
		return errors.New("book is already borrowed")
	}

	b.Available = false
	return nil
}

func (b *Book) Return() error {
	if b.Available {
		return errors.New("book is already returned")
	}

	b.Available = true
	return nil
}
