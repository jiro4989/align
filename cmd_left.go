package main

import (
	"github.com/nsf/termbox-go"
	"github.com/spf13/cobra"
)

func init() {
	RootCommand.AddCommand(leftCommand)
	leftCommand.Flags().StringP("pad", "p", " ", "Padding string")
	leftCommand.Flags().IntP("length", "n", -1, "Padding length")
	leftCommand.Flags().BoolP("write", "w", false, "Overwrite file")
	leftCommand.Flags().StringP("linefeed", "l", "\n", "Line feed")
	leftCommand.Flags().BoolP("termwidth", "t", false, "Terminal width")
}

var leftCommand = &cobra.Command{
	Use:     "left",
	Aliases: []string{"l"},
	Short:   "Align left command from file or stdin",
	Run: func(cmd *cobra.Command, args []string) {
		f := cmd.Flags()

		p, err := f.GetString("pad")
		if err != nil {
			panic(err)
		}

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

		useTermWidth, err := f.GetBool("termwidth")
		if err != nil {
			panic(err)
		}

		if useTermWidth {
			if err := termbox.Init(); err != nil {
				panic(err)
			}
			n, _ = termbox.Size()
			termbox.Close()
		}

		param := LogicHorizontalAlignParam{
			args:      args,
			pad:       p,
			length:    n,
			writeFile: writeFile,
			lineFeed:  lf,
			f:         AlignLeft,
		}
		if err := LogicHorizontalAlign(param); err != nil {
			panic(err)
		}
	},
}
