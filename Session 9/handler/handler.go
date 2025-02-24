package handler

import (
	"fmt"
	"os"
)

// ErrorHandler checks if the err is non-nil and, if so, prints
// the error message and exits the program.
func ErrorHandler(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
