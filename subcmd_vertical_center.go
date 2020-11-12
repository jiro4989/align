package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

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

		// 引数なしの場合は標準入力を処理
		if len(args) < 1 {
			args = readStdin()
			padded := AlignVerticalCenter(args, n)
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
			padded := AlignVerticalCenter(lines, n)

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
