package align

import (
	"strings"

	"github.com/mattn/go-runewidth"
)

// MaxStringWidth は最も長い見た目の文字列の長さを返す。
func MaxStringWidth(lines []string) (max int) {
	for _, v := range lines {
		l := runewidth.StringWidth(v)
		if max < l {
			max = l
		}
	}
	return max
}

// AlignLeft は文字列を左寄せする。
// 左寄せは見た目上の文字列の長さで揃える。
// length = -1のときは、引数文字列の最長の長さに合わせる。
// padは埋める文字列を指定する。埋める文字が見た目上でマルチバイトの場合は
// たとえlengthが奇数でも+1して偶数になるように調整する。
func AlignLeft(lines []string, length int, pad string) []string {
	if length == 0 || len(lines) < 1 {
		return lines
	}

	// 空白埋めする文字列がマルチバイト文字かどうか
	padWidtn := runewidth.StringWidth(pad)
	padIsMultiByteString := padWidtn == 2

	// -1のときは文字列の長さをalignの長さにする
	maxWidth := MaxStringWidth(lines)
	if length == -1 {
		length = maxWidth
	}

	// マルチバイト文字を使うときは長さを偶数に揃える
	if padIsMultiByteString && length%2 != 0 {
		length++
	}

	ret := []string{}
	for _, line := range lines {
		l := runewidth.StringWidth(line)
		diff := length - l
		if diff%2 != 0 {
			line += " "
			diff--
		}
		// Repeatするときにマルチバイト文字を使うときは2分の1にする
		if padIsMultiByteString {
			diff /= 2
		}
		s := line + strings.Repeat(pad, diff)
		ret = append(ret, s)
	}
	return ret
}

// AlignCenter は文字列を中央寄せする。
// 中央寄せは見た目上の文字列の長さで揃える。
// length = -1のときは、引数文字列の最長の長さに合わせる。
// padは埋める文字列を指定する。埋める文字が見た目上でマルチバイトの場合は
// たとえlengthが奇数でも+1して偶数になるように調整する。
func AlignCenter(lines []string, length int, pad string) []string {
	// TODO
	return nil
}

// AlignRight は文字列を右寄せする。
// 右寄せは見た目上の文字列の長さで揃える。
// length = -1のときは、引数文字列の最長の長さに合わせる。
// padは埋める文字列を指定する。埋める文字が見た目上でマルチバイトの場合は
// たとえlengthが奇数でも+1して偶数になるように調整する。
func AlignRight(lines []string, length int, pad string) []string {
	// TODO
	return nil
}
