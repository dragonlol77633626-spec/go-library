package storage

import (
	"encoding/csv"
	"encoding/json"
	"library-app/domain"
	"library-app/library"
	"os"
	"strconv"
)

type storageData struct {
	Books   []*domain.Book   `json:"books"`
	Readers []*domain.Reader `json:"readers"`
}

func SaveLibraryToJSON(filePath string, lib *library.Library) error {
	data := storageData{
		Books:   lib.Books,
		Readers: lib.Readers,
	}
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, jsonData, 0644)
}

func LoadLibraryFromJSON(filePath string) (*library.Library, error) {
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var data storageData
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return nil, err
	}
	lib := &library.Library{}
	lib.Books = data.Books
	lib.Readers = data.Readers
	lib.UpdateIDs()
	return lib, nil
}

func SaveBooksToCSV(filePath string, books []*domain.Book) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, book := range books {
		record := []string{
			strconv.Itoa(book.ID),
			book.Title,
			book.Author,
			strconv.Itoa(book.Year),
			strconv.FormatBool(book.IsIssued),
			strconv.FormatBool(book.IsReturned),
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}
	return nil
}

func LoadBooksFromCSV(filePath string) ([]*domain.Book, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var books []*domain.Book
	for _, record := range records {
		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, err
		}
		year, err := strconv.Atoi(record[3])
		if err != nil {
			return nil, err
		}
		isIssued, err := strconv.ParseBool(record[4])
		if err != nil {
			return nil, err
		}
		isReturned, err := strconv.ParseBool(record[5])
		if err != nil {
			return nil, err
		}

		book := &domain.Book{
			ID:         id,
			Title:      record[1],
			Author:     record[2],
			Year:       year,
			IsIssued:   isIssued,
			IsReturned: isReturned,
		}
		books = append(books, book)
	}
	return books, nil
}
