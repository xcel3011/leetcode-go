package string

import (
	"strings"
)

func replaceSpaces(S string, length int) string {
	b := strings.Builder{}
	for i := 0; i < length || i < len(S); i++ {
		if S[i] == ' ' {
			b.WriteString("%20")
			continue
		}
		b.Write([]byte{S[i]})
	}
	return b.String()
}
