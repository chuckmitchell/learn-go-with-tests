package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	Greet(os.Stdout, "charles")
}

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
