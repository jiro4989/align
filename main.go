package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime)
}

func main() {
	if err := Main(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}

func Main() error {
	return RootCommand.Execute()
}
