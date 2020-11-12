package main

import (
	"github.com/nsf/termbox-go"
	"github.com/spf13/cobra"
)

func init() {
	RootCommand.AddCommand(rightCommand)
	rightCommand.Flags().StringP("pad", "p", " ", "Padding string")
	rightCommand.Flags().IntP("length", "n", -1, "Padding length")
	rightCommand.Flags().BoolP("write", "w", false, "Overwrite file")
	rightCommand.Flags().StringP("linefeed", "l", "\n", "Line feed")
	rightCommand.Flags().BoolP("termwidth", "t", false, "Terminal width")
}

var rightCommand = &cobra.Command{
	Use:   "right",
	Short: "Align right command from file or stdin",
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
			f:         AlignRight,
		}
		LogicHorizontalAlign(param)
	},
}
