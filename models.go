package domain

import "fmt"

type Book struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	Year       int    `json:"year"`
	IsIssued   bool   `json:"is_issued"`
	IsReturned bool   `json:"is_returned"`
}

func (b Book) String() string {
	status := "available"
	if b.IsIssued {
		status = "issued"
	}
	return fmt.Sprintf("ID: %d | '%s' - %s (%d) (%s)", b.ID, b.Title, b.Author, b.Year, status)
}

type Reader struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
