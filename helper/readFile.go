package helper

import (
	"bufio"
	"fmt"
	"os"
)

// Reads the file line by line and stores it in an array
func ReadTextFile() []string {

	// Opens the file
	readTextFile, err := os.Open("books.txt")
	if err != nil {
		fmt.Println(err)
	}

	// Creates a scanner
	fileScanner := bufio.NewScanner(readTextFile)
	fileScanner.Split(bufio.ScanLines)

	var fileLines []string

	// Pushes each line to array
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readTextFile.Close()

	/* for _, v := range fileLines {
		fmt.Println(v)
	} */

	// fmt.Println(fileLines[0])

	return fileLines
}
