package models

import (
	"fmt"
	"time"
)

type User struct {
	Person
	JoinDate time.Time
}

func (u User) DisplayInfo() string {
	return fmt.Sprintf(
		"%s | Join Date: %s",
		u.Person.DisplayInfo(), u.JoinDate.Format("2006-01-02"),
	)
}

func (u User) GetBorrowedBooksCount() int {
	return 0
}
