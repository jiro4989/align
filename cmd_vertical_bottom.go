package main

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCommand.AddCommand(verticalBottomCommand)
	verticalBottomCommand.Flags().IntP("length", "n", -1, "Padding length")
	verticalBottomCommand.Flags().BoolP("write", "w", false, "Overwrite file")
	verticalBottomCommand.Flags().StringP("linefeed", "l", "\n", "Line feed")
}

var verticalBottomCommand = &cobra.Command{
	Use:     "vertical-bottom",
	Aliases: []string{"vb"},
	Short:   "Align vertical bottom command from file or stdin",
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
			f:         AlignVerticalBottom,
		}
		if err := LogicVerticalAlign(param); err != nil {
			panic(err)
		}
	},
}
