package log

import (
	"fmt"
	"os"
)

func Fatal(a ...interface{}) {
	fmt.Fprintf(os.Stderr, "Fatal: ")
	fmt.Fprintln(os.Stderr, a)
	os.Exit(1)
}

func Error(a ...interface{}) {
	fmt.Fprintf(os.Stderr, "Error: ")
	fmt.Fprintln(os.Stderr, a)
}
