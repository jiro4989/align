package main

import (
	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize()
}

var RootCommand = &cobra.Command{
	Use:     "align",
	Short:   "align is text align command",
	Example: "align right README.md",
	Version: Version,
	Long: `align is text align command.
align [left|center|right]-justify text files.
or does stdin too.

Repository: https://github.com/jiro4989/align
    Author: jiro4989
	`,
}
