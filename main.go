package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jiro4989/align/subcmd"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime)
	subcmd.Version = Version
}

func main() {
	if err := subcmd.RootCommand.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}
