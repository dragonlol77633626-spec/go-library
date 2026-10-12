package main

import (
	"fmt"
	"library-app/library"
	"library-app/storage"
	"os"
)

func main() {
	lib := &library.Library{}

	_, _ = lib.AddBook("The Hobbit", "J.R.R. Tolkien", 1937)
	_, _ = lib.AddBook("1984", "George Orwell", 1949)

	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . <command> <filename.csv>")
		return
	}

	command := os.Args[1]
	fileName := os.Args[2]

	switch command {
	case "save":
		err := storage.SaveBooksToCSV(fileName, lib.Books)
		if err != nil {
			fmt.Println("Save CSV error:", err)
			return
		}
		fmt.Printf("Success: Books saved to %s\n", fileName)

	case "load":
		books, err := storage.LoadBooksFromCSV(fileName)
		if err != nil {
			fmt.Println("Load CSV error:", err)
			return
		}
		lib.Books = books
		lib.UpdateIDs()
		fmt.Printf("Success: Books loaded from %s\n", fileName)
		lib.ListAllBooks()

	default:
		fmt.Println("Unknown command. Use 'save' or 'load'")
	}
}
