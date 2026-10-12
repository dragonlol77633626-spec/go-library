package library

import (
	"errors"
	"fmt"
	"library-app/domain"
)

type Library struct {
	Books        []*domain.Book   `json:"books"`
	Readers      []*domain.Reader `json:"readers"`
	lastBookID   int
	lastReaderID int
}

func (lib *Library) UpdateIDs() {
	maxBookID := 0
	for _, book := range lib.Books {
		if book.ID > maxBookID {
			maxBookID = book.ID
		}
	}
	lib.lastBookID = maxBookID

	maxReaderID := 0
	for _, reader := range lib.Readers {
		if reader.ID > maxReaderID {
			maxReaderID = reader.ID
		}
	}
	lib.lastReaderID = maxReaderID
}

func (lib *Library) AddBook(title, author string, year int) (*domain.Book, error) {
	for _, b := range lib.Books {
		if b.Title == title && b.Author == author && b.Year == year {
			return nil, errors.New("book with this title, author and year already exists")
		}
	}
	lib.lastBookID++
	newBook := &domain.Book{
		ID:         lib.lastBookID,
		Title:      title,
		Author:     author,
		Year:       year,
		IsReturned: true,
	}
	lib.Books = append(lib.Books, newBook)
	return newBook, nil
}

func (lib *Library) AddReader(name string) *domain.Reader {
	lib.lastReaderID++
	newReader := &domain.Reader{
		ID:   lib.lastReaderID,
		Name: name,
	}
	lib.Readers = append(lib.Readers, newReader)
	return newReader
}

func (lib *Library) FindBookByID(id int) (*domain.Book, error) {
	for _, book := range lib.Books {
		if book.ID == id {
			return book, nil
		}
	}
	return nil, errors.New("book not found")
}

func (lib *Library) FindReaderByID(id int) (*domain.Reader, error) {
	for _, reader := range lib.Readers {
		if reader.ID == id {
			return reader, nil
		}
	}
	return nil, errors.New("reader not found")
}

func (lib *Library) IssueBookToReader(bookID int, readerID int) error {
	book, err := lib.FindBookByID(bookID)
	if err != nil {
		return err
	}
	_, err = lib.FindReaderByID(readerID)
	if err != nil {
		return err
	}
	if book.IsIssued {
		return errors.New("book is already issued")
	}
	book.IsIssued = true
	book.IsReturned = false
	return nil
}

func (lib *Library) ReturnBook(bookID int) error {
	book, err := lib.FindBookByID(bookID)
	if err != nil {
		return err
	}
	if book.IsReturned || !book.IsIssued {
		return errors.New("book is already in the library")
	}
	book.IsIssued = false
	book.IsReturned = true
	return nil
}

func (lib *Library) ListAllBooks() {
	for _, book := range lib.Books {
		fmt.Println(book)
	}
}
