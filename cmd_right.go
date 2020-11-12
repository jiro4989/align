package main

import (
	"github.com/nsf/termbox-go"
	"github.com/spf13/cobra"
)

var (
	rightPad       string
	rightLength    int
	rightWrite     bool
	rightLineFeed  string
	rightTermWidth bool
)

func init() {
	RootCommand.AddCommand(rightCommand)
	rightCommand.Flags().StringVarP(&rightPad, "pad", "p", " ", "Padding string")
	rightCommand.Flags().IntVarP(&rightLength, "length", "n", -1, "Padding length")
	rightCommand.Flags().BoolVarP(&rightWrite, "write", "w", false, "Overwrite file")
	rightCommand.Flags().StringVarP(&rightLineFeed, "linefeed", "l", "\n", "Line feed")
	rightCommand.Flags().BoolVarP(&rightTermWidth, "termwidth", "t", false, "Terminal width")
}

var rightCommand = &cobra.Command{
	Use:     "right",
	Aliases: []string{"r"},
	Short:   "Align right command from file or stdin",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			p            = rightPad
			n            = rightLength
			writeFile    = rightWrite
			lf           = rightLineFeed
			useTermWidth = rightTermWidth
		)

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
			f:         AlignRight,
		}
		if err := LogicHorizontalAlign(param); err != nil {
			panic(err)
		}
	},
}
