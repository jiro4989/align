package main

import (
	"github.com/nsf/termbox-go"
	"github.com/spf13/cobra"
)

var (
	leftPad       string
	leftLength    int
	leftWrite     bool
	leftLineFeed  string
	leftTermWidth bool
)

func init() {
	RootCommand.AddCommand(leftCommand)
	leftCommand.Flags().StringVarP(&leftPad, "pad", "p", " ", "Padding string")
	leftCommand.Flags().IntVarP(&leftLength, "length", "n", -1, "Padding length")
	leftCommand.Flags().BoolVarP(&leftWrite, "write", "w", false, "Overwrite file")
	leftCommand.Flags().StringVarP(&leftLineFeed, "linefeed", "l", "\n", "Line feed")
	leftCommand.Flags().BoolVarP(&leftTermWidth, "termwidth", "t", false, "Terminal width")
}

var leftCommand = &cobra.Command{
	Use:     "left",
	Aliases: []string{"l"},
	Short:   "Align left command from file or stdin",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			p            = leftPad
			n            = leftLength
			writeFile    = leftWrite
			lf           = leftLineFeed
			useTermWidth = leftTermWidth
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
			f:         AlignLeft,
		}
		if err := LogicHorizontalAlign(param); err != nil {
			panic(err)
		}
	},
}
