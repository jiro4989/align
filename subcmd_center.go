package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jiro4989/align/align"
	"github.com/nsf/termbox-go"
	"github.com/spf13/cobra"
)

func init() {
	RootCommand.AddCommand(centerCommand)
	centerCommand.Flags().StringP("pad", "p", " ", "Padding string")
	centerCommand.Flags().IntP("length", "n", -1, "Padding length")
	centerCommand.Flags().BoolP("write", "w", false, "Overwrite file")
	centerCommand.Flags().StringP("linefeed", "l", "\n", "Line feed")
	centerCommand.Flags().BoolP("termwidth", "t", false, "Terminal width")
}

var centerCommand = &cobra.Command{
	Use:   "center",
	Short: "Align center command from file or stdin",
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

		// 引数なしの場合は標準入力を処理
		if len(args) < 1 {
			args = readStdin()
			padded := align.AlignCenter(args, n, p)
			for _, v := range padded {
				fmt.Println(v)
			}
			return
		}

		for _, fn := range args {
			b, err := ioutil.ReadFile(fn)
			if err != nil {
				panic(err)
			}
			s := string(b)
			lines := strings.Split(s, lf)
			padded := align.AlignCenter(lines, n, p)

			// ファイル上書き指定があれば上書き
			if writeFile {
				b := []byte(strings.Join(padded, lf))
				if err := ioutil.WriteFile(fn, b, os.ModePerm); err != nil {
					panic(err)
				}
				continue
			}
			for _, v := range padded {
				fmt.Println(v)
			}
		}
	},
}
