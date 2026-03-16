package models

import (
	"fmt"
	"time"
)

type BorrowRecord struct {
	ID         int
	UserID     int
	BookID     int
	BorrowDate time.Time
	ReturnDate *time.Time
	Status     string
}

func (br BorrowRecord) DisplayInfo() string {
	rd := "Not returned yet"

	if br.ReturnDate != nil {
		rd = br.ReturnDate.Format("2006-01-02")
	}

	return fmt.Sprintf(
		"ID: %d | UserID: %d | BookID: %d | BorrowDate: %s | ReturnDate: %s | Status: %s",
		br.ID, br.UserID, br.BookID, br.BorrowDate, rd, br.Status,
	)
}

func (br *BorrowRecord) MarkAsReturned() error {
	now := time.Now()

	br.Status = "returned"
	br.ReturnDate = &now

	return nil
}
