package main

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCommand.AddCommand(verticalCenterCommand)
	verticalCenterCommand.Flags().IntP("length", "n", -1, "Padding length")
	verticalCenterCommand.Flags().BoolP("write", "w", false, "Overwrite file")
	verticalCenterCommand.Flags().StringP("linefeed", "l", "\n", "Line feed")
}

var verticalCenterCommand = &cobra.Command{
	Use:     "vertical-center",
	Aliases: []string{"vc"},
	Short:   "Align vertical center command from file or stdin",
	Run: func(cmd *cobra.Command, args []string) {
		f := cmd.Flags()

		n, err := f.GetInt("length")
		if err != nil {
			panic(err)
		}

		writeFile, err := f.GetBool("write")
		if err != nil {
			panic(err)
		}

		lf, err := f.GetString("linefeed")
		if err != nil {
			panic(err)
		}

		param := LogicVerticalAlignParam{
			args:      args,
			length:    n,
			writeFile: writeFile,
			lineFeed:  lf,
			f:         AlignVerticalCenter,
		}
		if err := LogicVerticalAlign(param); err != nil {
			panic(err)
		}
	},
}
