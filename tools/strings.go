package tools

import (
	"math/rand"
	"strings"
)

func ReplaceStringOnIndex(s, old, replace string, n int) string {
	i := 0
	for m := 1; m <= n; m++ {
		x := strings.Index(s[i:], old)
		if x < 0 {
			break
		}
		i += x
		if m == n {
			return s[:i] + replace + s[i+len(old):]
		}
		i += len(old)
	}
	return s
}

func CleanString(subject string, findme string, replace string) string {
	return strings.Replace(subject, findme, replace, -1)
}

var randomRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = randomRunes[rand.Intn(len(randomRunes))]
	}
	return string(b)
}

func SplitWithoutEmpty(data string, separator string) []string {
	raws := strings.Split(data, separator)
	var words []string
	for _, item := range raws {
		if item != "" {
			words = append(words, item)
		}
	}
	return words
}
