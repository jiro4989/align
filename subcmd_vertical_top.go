package main

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCommand.AddCommand(verticalTopCommand)
	verticalTopCommand.Flags().IntP("length", "n", -1, "Padding length")
	verticalTopCommand.Flags().BoolP("write", "w", false, "Overwrite file")
	verticalTopCommand.Flags().StringP("linefeed", "l", "\n", "Line feed")
}

var verticalTopCommand = &cobra.Command{
	Use:     "vertical-top",
	Aliases: []string{"vt"},
	Short:   "Align vertical top command from file or stdin",
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
			f:         AlignVerticalTop,
		}
		LogicVerticalAlign(param)
	},
}
