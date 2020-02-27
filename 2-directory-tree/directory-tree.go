package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	currentWorkingDirectory, _ := os.Getwd() // Get the current working directory
	scanDirectory(currentWorkingDirectory, 0) // Start searching the current working directory
}

/**
 * Search the given directory for all files
 */
func scanDirectory(directory string, level int) {
	currentDirectoryListing, _ := ioutil.ReadDir(directory)
	levelIncrement := level + 1

	for _, element := range currentDirectoryListing {
		fmt.Print(strings.Repeat("  ", level))

		// Is the element a directory?
		if element.IsDir() {
			// Output the current folder
			fmt.Print( "📁 ", element.Name(), "\n")
			// Then go and rerun the for loop within.
			scanDirectory(directory + "/" + element.Name(), levelIncrement)
		} else {
			// Output the file
			fmt.Print("📄 ", element.Name(), "\n")
		}
	}
}