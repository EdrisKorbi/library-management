package library

import (
	"errors"
	"time"

	"github.com/EdrisKorbi/library-management/models"
)

type Library struct {
	books         []models.Book
	users         []models.User
	borrowRecords []models.BorrowRecord
	nextBookID    int
	nextUserID    int
	nextRecordID  int
}

const (
	StatusBorrowed = "borrowed"
	StatusReturned = "returned"
)

func NewLibrary() *Library {
	return &Library{
		books:         []models.Book{},
		users:         []models.User{},
		borrowRecords: []models.BorrowRecord{},
		nextBookID:    1,
		nextUserID:    1,
		nextRecordID:  1,
	}
}

func (lib *Library) AddUser(name, email string) models.User {
	user := models.User{
		Person: models.Person{
			ID:    lib.nextUserID,
			Name:  name,
			Email: email,
		},
		JoinDate: time.Now(),
	}

	lib.users = append(lib.users, user)
	lib.nextUserID++

	return user

}

func (lib *Library) AddBook(title, author, isbn string) models.Book {
	book := models.Book{
		ID:        lib.nextBookID,
		Title:     title,
		Author:    author,
		ISBN:      isbn,
		Available: true,
	}

	lib.books = append(lib.books, book)
	lib.nextBookID++

	return book
}

func (lib *Library) FindBookByID(id int) *models.Book {
	for i := range lib.books {
		if lib.books[i].ID == id {
			return &lib.books[i]
		}
	}
	return nil
}

func (lib *Library) FindUserByID(id int) *models.User {
	for i := range lib.users {
		if lib.users[i].ID == id {
			return &lib.users[i]
		}
	}
	return nil
}

func (lib *Library) ListBooks() []models.Book {
	return lib.books
}

func (lib *Library) ListUsers() []models.User {
	return lib.users
}

func (lib *Library) ListBorrowHistory() []models.BorrowRecord {
	return lib.borrowRecords
}

func (lib *Library) BorrowBook(userID, bookID int) error {
	user := lib.FindUserByID(userID)
	if user == nil {
		return errors.New("user not found")
	}

	book := lib.FindBookByID(bookID)
	if book == nil {
		return errors.New("book not found")
	}

	if err := book.Borrow(); err != nil {
		return err
	}

	br := models.BorrowRecord{
		ID:         lib.nextRecordID,
		UserID:     userID,
		BookID:     bookID,
		BorrowDate: time.Now(),
		ReturnDate: nil,
		Status:     StatusBorrowed,
	}

	lib.borrowRecords = append(lib.borrowRecords, br)
	lib.nextRecordID++

	return nil
}

func (lib *Library) ReturnBook(userID, bookID int) error {
	book := lib.FindBookByID(bookID)
	if book == nil {
		return errors.New("book not found")
	}

	user := lib.FindUserByID(userID)
	if user == nil {
		return errors.New("user not found")
	}
	recordIndex := -1
	for i := range lib.borrowRecords {
		br := lib.borrowRecords[i]
		if br.UserID == userID && br.BookID == bookID && br.Status == StatusBorrowed {
			recordIndex = i
			break
		}
	}

	if recordIndex == -1 {
		return errors.New("borrow record not found")
	}

	if err := book.Return(); err != nil {
		return err
	}

	if err := lib.borrowRecords[recordIndex].MarkAsReturned(); err != nil {
		return err
	}

	return nil
}

func (lib *Library) GetUserBorrowedRecords(userID int) []models.Book {
	var borrowedBooks []models.Book

	for i := range lib.borrowRecords {
		br := lib.borrowRecords[i]
		if br.UserID == userID && br.Status == StatusBorrowed {
			book := lib.FindBookByID(br.BookID)

			if book != nil {
				borrowedBooks = append(borrowedBooks, *book)
			}
		}
	}

	return borrowedBooks
}

func (lib *Library) IsBookAvailable(bookID int) bool {
	book := lib.FindBookByID(bookID)

	return book != nil && book.Available
}
