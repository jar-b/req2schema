package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("unexpected number of args")
	}
	b, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(fmt.Errorf("reading input: %w", err))
	}

	out, err := requestToSchema(b)
	if err != nil {
		log.Fatal(fmt.Errorf("schema generation: %w", err))
	}
	fmt.Println(string(out))
}
