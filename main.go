package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-1-atakansenoglu/helper"
)

var Args []string = os.Args[1:]
var BooksArray []string = helper.ReadTextFile()

func main() {

	switch checkConsoleInput() {
	case "search":
		searchBookshelf(getBookName())
	case "list":
		listBooks(BooksArray)
	case "":
		fmt.Println("Kitap ismi yazmadınız!")
	case "err":
		fmt.Println("Yanlış veya eksik komut girdiniz!")
	}
}

// Check console input
func checkConsoleInput() string {

	if len(Args) <= 0 {
		fmt.Printf("search => arama işlemi için \n list => listeleme işlemi için\n")
		return ""

	} else if len(Args) > 0 && Args[0] == "search" {
		if len(Args) == 1 {
			return ""
		}
		return "search"

	} else if len(Args) > 0 && Args[0] == "list" {
		return "list"

	} else {
		return "err"
	}

}

// Get book name
func getBookName() string {

	// Creates a slice of length zero to store console input as a string.
	var searchedBookName = make([]string, 0)

	//Pushes console input into the slice.
	searchedBookName = append(searchedBookName, Args...)

	//Concatenates the slice starting from second index and returns a string.
	searchedBookNameString := strings.Join(searchedBookName[1:], " ")
	fmt.Println("Aranan kitap:", searchedBookNameString)

	return searchedBookNameString
}

// Search for a book
// 0. sıradaki kitabı bulamıyor!!
func searchBookshelf(bookName string) {

	if contains(BooksArray, bookName) {
	} else {
		fmt.Println("Kitap bulunamadı.")
	}

}

// List all books
func listBooks(arr []string) {

	fmt.Println("Listed all books:")

	for i, v := range arr {
		fmt.Printf("%d: %s\n", i, v)
	}

}

// Contains
func contains(arr []string, str string) bool {
	for i, v := range arr {
		if v == str {
			fmt.Printf("Kitap: %s\nSıra: %d\n", v, i+1)
			return true
		}
	}
	return false
}
