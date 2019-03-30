package align

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
}
