package models

import "fmt"

type Person struct {
	ID    int
	Name  string
	Email string
}

func (p Person) DisplayInfo() string {
	return fmt.Sprintf("ID: %d | Name: %s | Email: %s", p.ID, p.Name, p.Email)
}
