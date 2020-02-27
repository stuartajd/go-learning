package main

import (
	"flag" // Used to parse the arguments
	"fmt" // Used to output
	"strconv" // Used to convert the flags
)

func main() {
	// Set the current variable to be 0
	var result int = 0

	// Parse all the arguments on the command
	flag.Parse()

	// Loop over all the arguments
	for i := 0; i < len(flag.Args()); i++ {
		// Convert to an integer, flag.Arg() returns a string
		value,err := strconv.Atoi(flag.Arg(i))

		// Check if there is an error with the conversion, if not add to our result
		if err == nil {
			result += value
		}
	}

	// When we've looped over, just print out the result
	fmt.Println(result)
}