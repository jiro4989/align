package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

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

		// 引数なしの場合は標準入力を処理
		if len(args) < 1 {
			args = readStdin()
			padded := AlignVerticalBottom(args, n)
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
			padded := AlignVerticalBottom(lines, n)

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
