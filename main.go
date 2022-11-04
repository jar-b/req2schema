package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var out string

func init() {
	// slightly better usage output
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [flags] [filename]\n\nFlags:\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	flag.StringVar(&out, "out", "", "output file (optional, prints to stdout if omitted)")
	flag.Parse()

	if flag.NArg() != 1 {
		log.Fatal("unexpected number of args")
	}
	in := flag.Arg(0)

	b, err := os.ReadFile(in)
	if err != nil {
		log.Fatal(fmt.Errorf("reading input: %w", err))
	}

	schm, err := requestToSchema(b)
	if err != nil {
		log.Fatal(fmt.Errorf("schema generation: %w", err))
	}

	if out != "" {
		fmt.Printf("writing to: %s\n", out)
		os.WriteFile(out, schm, 0644)
	} else {
		fmt.Println(string(schm))
	}
}
