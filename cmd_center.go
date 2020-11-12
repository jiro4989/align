package main

import (
	"github.com/nsf/termbox-go"
	"github.com/spf13/cobra"
)

var (
	centerPad       string
	centerLength    int
	centerWrite     bool
	centerLineFeed  string
	centerTermWidth bool
)

func init() {
	RootCommand.AddCommand(centerCommand)
	centerCommand.Flags().StringVarP(&centerPad, "pad", "p", " ", "Padding string")
	centerCommand.Flags().IntVarP(&centerLength, "length", "n", -1, "Padding length")
	centerCommand.Flags().BoolVarP(&centerWrite, "write", "w", false, "Overwrite file")
	centerCommand.Flags().StringVarP(&centerLineFeed, "linefeed", "l", "\n", "Line feed")
	centerCommand.Flags().BoolVarP(&centerTermWidth, "termwidth", "t", false, "Terminal width")
}

var centerCommand = &cobra.Command{
	Use:     "center",
	Aliases: []string{"c"},
	Short:   "Align center command from file or stdin",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			p            = centerPad
			n            = centerLength
			writeFile    = centerWrite
			lf           = centerLineFeed
			useTermWidth = centerTermWidth
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
			f:         AlignCenter,
		}
		if err := LogicHorizontalAlign(param); err != nil {
			panic(err)
		}
	},
}
