package utils

import (
	"fmt"
	"os"
)

// dd function that prints the argument and exits the program
func Dump(v interface{}) {
	fmt.Printf("%+v\n", v)
	os.Exit(1)
}
