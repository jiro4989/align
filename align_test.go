package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxStringWidth(t *testing.T) {
	assert.Equal(t, 0, MaxStringWidth([]string{}))
	assert.Equal(t, 0, MaxStringWidth([]string{""}))
	assert.Equal(t, 1, MaxStringWidth([]string{"1"}))
	assert.Equal(t, 2, MaxStringWidth([]string{"12"}))
	assert.Equal(t, 2, MaxStringWidth([]string{"あ"}))
	assert.Equal(t, 2, MaxStringWidth([]string{"あ", "1"}))
	assert.Equal(t, 3, MaxStringWidth([]string{"あa", "1"}))
	assert.Equal(t, 3, MaxStringWidth([]string{"abc", "1"}))
}

func TestAlignLeft(t *testing.T) {
	assert.Equal(t, []string{}, AlignLeft([]string{}, -1, " "))
	assert.Equal(t, []string{}, AlignLeft([]string{}, -1, "　"))
	assert.Equal(t, []string{}, AlignLeft([]string{}, 0, " "))
	assert.Equal(t, []string{}, AlignLeft([]string{}, 0, "　"))
	assert.Equal(t, []string{}, AlignLeft([]string{}, 1, " "))
	assert.Equal(t, []string{}, AlignLeft([]string{}, 1, "　"))

	assert.Equal(t, []string{"a"}, AlignLeft([]string{"a"}, -1, " "))
	assert.Equal(t, []string{"a "}, AlignLeft([]string{"a"}, -1, "　"))
	assert.Equal(t, []string{"a"}, AlignLeft([]string{"a"}, 0, " "))
	assert.Equal(t, []string{"a"}, AlignLeft([]string{"a"}, 0, "　"))
	assert.Equal(t, []string{"a"}, AlignLeft([]string{"a"}, 1, " "))
	assert.Equal(t, []string{"a "}, AlignLeft([]string{"a"}, 1, "　"))
	assert.Equal(t, []string{"a "}, AlignLeft([]string{"a"}, 2, " "))
	assert.Equal(t, []string{"a "}, AlignLeft([]string{"a"}, 2, "　"))
	assert.Equal(t, []string{"a  "}, AlignLeft([]string{"a"}, 3, " "))
	assert.Equal(t, []string{"a 　"}, AlignLeft([]string{"a"}, 3, "　"))
	assert.Equal(t, []string{"a   "}, AlignLeft([]string{"a"}, 4, " "))
	assert.Equal(t, []string{"a 　"}, AlignLeft([]string{"a"}, 4, "　"))

	assert.Equal(t, []string{"a ", "あ"}, AlignLeft([]string{"a", "あ"}, -1, " "))
	assert.Equal(t, []string{"a", "a"}, AlignLeft([]string{"a", "a"}, -1, " "))
	assert.Equal(t, []string{"a    ", "12345"}, AlignLeft([]string{"a", "12345"}, -1, " "))
	assert.Equal(t, []string{"a     ", "123456"}, AlignLeft([]string{"a", "123456"}, -1, " "))
	assert.Equal(t, []string{"a ", "a "}, AlignLeft([]string{"a", "a"}, 2, " "))
	assert.Equal(t, []string{"a  ", "a  "}, AlignLeft([]string{"a", "a"}, 3, " "))
	assert.Equal(t, []string{"a ", "a "}, AlignLeft([]string{"a", "a"}, -1, "　"))
	assert.Equal(t, []string{"a 　　", "12345 "}, AlignLeft([]string{"a", "12345"}, -1, "　"))
	assert.Equal(t, []string{"a 　　", "123456"}, AlignLeft([]string{"a", "123456"}, -1, "　"))
	assert.Equal(t, []string{"a ", "a "}, AlignLeft([]string{"a", "a"}, 2, "　"))
	assert.Equal(t, []string{"a 　", "a 　"}, AlignLeft([]string{"a", "a"}, 3, "　"))

	assert.Equal(t, []string{"123456", "a     "}, AlignLeft([]string{"123456", "a"}, 4, " "))
}

func TestAlignCenter(t *testing.T) {
	assert.Equal(t, []string{}, AlignCenter([]string{}, -1, " "))
	assert.Equal(t, []string{}, AlignCenter([]string{}, -1, "　"))
	assert.Equal(t, []string{}, AlignCenter([]string{}, 0, " "))
	assert.Equal(t, []string{}, AlignCenter([]string{}, 0, "　"))
	assert.Equal(t, []string{}, AlignCenter([]string{}, 1, " "))
	assert.Equal(t, []string{}, AlignCenter([]string{}, 1, "　"))

	assert.Equal(t, []string{"a"}, AlignCenter([]string{"a"}, -1, " "))
	assert.Equal(t, []string{"a"}, AlignCenter([]string{"a"}, 0, " "))
	assert.Equal(t, []string{"a"}, AlignCenter([]string{"a"}, 1, " "))
	assert.Equal(t, []string{"a "}, AlignCenter([]string{"a"}, 2, " "))
	assert.Equal(t, []string{" a "}, AlignCenter([]string{"a"}, 3, " "))
	assert.Equal(t, []string{" a  "}, AlignCenter([]string{"a"}, 4, " "))

	assert.Equal(t, []string{"a"}, AlignCenter([]string{"a"}, -1, "　"))
	assert.Equal(t, []string{"a"}, AlignCenter([]string{"a"}, 0, "　"))
	assert.Equal(t, []string{"a"}, AlignCenter([]string{"a"}, 1, "　"))
	assert.Equal(t, []string{"a "}, AlignCenter([]string{"a"}, 2, "　"))
	assert.Equal(t, []string{" a "}, AlignCenter([]string{"a"}, 3, "　"))
	assert.Equal(t, []string{" a  "}, AlignCenter([]string{"a"}, 4, "　"))
	assert.Equal(t, []string{"　a　"}, AlignCenter([]string{"a"}, 5, "　"))
	assert.Equal(t, []string{"　a 　"}, AlignCenter([]string{"a"}, 6, "　"))
	assert.Equal(t, []string{"　 a  　"}, AlignCenter([]string{"a"}, 8, "　"))
	assert.Equal(t, []string{"　　a　　"}, AlignCenter([]string{"a"}, 9, "　"))
	assert.Equal(t, []string{"　　a 　　"}, AlignCenter([]string{"a"}, 10, "　"))

	assert.Equal(t, []string{"　a　", "12345"}, AlignCenter([]string{"a", "12345"}, -1, "　"))

	assert.Equal(t, []string{"123456", "  a   "}, AlignCenter([]string{"123456", "a"}, 4, " "))
}

func TestAlignRight(t *testing.T) {
	assert.Equal(t, []string{}, AlignRight([]string{}, -1, " "))
	assert.Equal(t, []string{}, AlignRight([]string{}, -1, "　"))
	assert.Equal(t, []string{}, AlignRight([]string{}, 0, " "))
	assert.Equal(t, []string{}, AlignRight([]string{}, 0, "　"))
	assert.Equal(t, []string{}, AlignRight([]string{}, 1, " "))
	assert.Equal(t, []string{}, AlignRight([]string{}, 1, "　"))

	assert.Equal(t, []string{"a"}, AlignRight([]string{"a"}, -1, " "))
	assert.Equal(t, []string{" a"}, AlignRight([]string{"a"}, -1, "　"))
	assert.Equal(t, []string{"a"}, AlignRight([]string{"a"}, 0, " "))
	assert.Equal(t, []string{"a"}, AlignRight([]string{"a"}, 0, "　"))
	assert.Equal(t, []string{"a"}, AlignRight([]string{"a"}, 1, " "))
	assert.Equal(t, []string{" a"}, AlignRight([]string{"a"}, 1, "　"))
	assert.Equal(t, []string{" a"}, AlignRight([]string{"a"}, 2, " "))
	assert.Equal(t, []string{" a"}, AlignRight([]string{"a"}, 2, "　"))
	assert.Equal(t, []string{"  a"}, AlignRight([]string{"a"}, 3, " "))
	assert.Equal(t, []string{"　 a"}, AlignRight([]string{"a"}, 3, "　"))
	assert.Equal(t, []string{"   a"}, AlignRight([]string{"a"}, 4, " "))
	assert.Equal(t, []string{"　 a"}, AlignRight([]string{"a"}, 4, "　"))

	assert.Equal(t, []string{" a", "あ"}, AlignRight([]string{"a", "あ"}, -1, " "))
	assert.Equal(t, []string{"a", "a"}, AlignRight([]string{"a", "a"}, -1, " "))
	assert.Equal(t, []string{"    a", "12345"}, AlignRight([]string{"a", "12345"}, -1, " "))
	assert.Equal(t, []string{"     a", "123456"}, AlignRight([]string{"a", "123456"}, -1, " "))
	assert.Equal(t, []string{" a", " a"}, AlignRight([]string{"a", "a"}, 2, " "))
	assert.Equal(t, []string{"  a", "  a"}, AlignRight([]string{"a", "a"}, 3, " "))
	assert.Equal(t, []string{" a", " a"}, AlignRight([]string{"a", "a"}, -1, "　"))
	assert.Equal(t, []string{"　　 a", " 12345"}, AlignRight([]string{"a", "12345"}, -1, "　"))
	assert.Equal(t, []string{"　　 a", "123456"}, AlignRight([]string{"a", "123456"}, -1, "　"))
	assert.Equal(t, []string{" a", " a"}, AlignRight([]string{"a", "a"}, 2, "　"))
	assert.Equal(t, []string{"　 a", "　 a"}, AlignRight([]string{"a", "a"}, 3, "　"))

	assert.Equal(t, []string{"123456", "     a"}, AlignRight([]string{"123456", "a"}, 4, " "))
}

func TestAlignVerticalTop(t *testing.T) {
	tests := []struct {
		desc   string
		lines  []string
		height int
		want   []string
	}{
		{
			desc: "正常系: 1行上に揃える",
			lines: []string{
				"hello",
				"world",
			},
			height: 3,
			want: []string{
				"hello",
				"world",
				"",
			},
		},
		{
			desc: "正常系: 変更なし",
			lines: []string{
				"hello",
				"world",
			},
			height: 2,
			want: []string{
				"hello",
				"world",
			},
		},
		{
			desc: "正常系: heightのほうがlinesより小さい場合は変更なし",
			lines: []string{
				"hello",
				"world",
			},
			height: 1,
			want: []string{
				"hello",
				"world",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)
			got := AlignVerticalTop(tt.lines, tt.height)
			assert.Equal(tt.want, got)
		})
	}
}
