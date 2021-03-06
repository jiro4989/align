package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type LogicHorizontalAlignParam struct {
	args      []string
	pad       string
	length    int
	writeFile bool
	lineFeed  string
	f         func([]string, int, string) []string
}

func LogicHorizontalAlign(param LogicHorizontalAlignParam) error {
	var (
		args      = param.args
		n         = param.length
		p         = param.pad
		lf        = param.lineFeed
		writeFile = param.writeFile
		f         = param.f
	)

	// 引数なしの場合は標準入力を処理
	if len(args) < 1 {
		args = readStdin()
		padded := f(args, n, p)
		for _, v := range padded {
			fmt.Println(v)
		}
		return nil
	}

	for _, fn := range args {
		b, err := ioutil.ReadFile(fn)
		if err != nil {
			return err
		}
		s := string(b)
		lines := strings.Split(s, lf)
		padded := f(lines, n, p)

		// ファイル上書き指定があれば上書き
		if writeFile {
			b := []byte(strings.Join(padded, lf))
			if err := ioutil.WriteFile(fn, b, os.ModePerm); err != nil {
				return err
			}
			continue
		}
		for _, v := range padded {
			fmt.Println(v)
		}
	}
	return nil
}

type LogicVerticalAlignParam struct {
	args      []string
	length    int
	writeFile bool
	lineFeed  string
	f         func([]string, int) []string
}

func LogicVerticalAlign(param LogicVerticalAlignParam) error {
	var (
		args      = param.args
		n         = param.length
		lf        = param.lineFeed
		writeFile = param.writeFile
		f         = param.f
	)

	// 引数なしの場合は標準入力を処理
	if len(args) < 1 {
		args = readStdin()
		padded := f(args, n)
		for _, v := range padded {
			fmt.Println(v)
		}
		return nil
	}

	for _, fn := range args {
		b, err := ioutil.ReadFile(fn)
		if err != nil {
			return err
		}
		s := string(b)
		lines := strings.Split(s, lf)
		padded := f(lines, n)

		// ファイル上書き指定があれば上書き
		if writeFile {
			b := []byte(strings.Join(padded, lf))
			if err := ioutil.WriteFile(fn, b, os.ModePerm); err != nil {
				return err
			}
			continue
		}
		for _, v := range padded {
			fmt.Println(v)
		}
	}
	return nil
}
