package strings

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func Basename(s string) string {
	slash := strings.LastIndex(s, string(os.PathSeparator))
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func IntsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func Join1(values []string, sep string) string {
	var buf strings.Builder
	for i, v := range values {
		if i > 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(v)
	}
	return buf.String()
}

func Join2(values []string, sep string) string {
	var buf bytes.Buffer
	for i, v := range values {
		if i > 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(v)
	}
	return buf.String()
}

// IsPalindrome reports whether s reads the same forward and backward.
// 但是这个实现的问题在于采用的byte序而不是rune序列，导致对于非ASCII字符的处理不正确
func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

// 支持忽略空格，特殊字符，支持非ASCII字符
func IsPalindrome2(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
