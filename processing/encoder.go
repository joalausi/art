package processing

import (
	"fmt"
	"strings"
)

func encodeLine(line string) string {
	if line == "" {
		return ""
	}

	var result strings.Builder
	i := 0
	n := len(line)

	for i < n {
		maxPatLen := 3 // You can tweak this (2 or 3)
		found := false

		// Try to match patterns of length 3, 2, 1 (in that order)
		for patLen := maxPatLen; patLen >= 1; patLen-- {
			if i+patLen > n {
				continue
			}
			pattern := line[i : i+patLen]
			count := 1
			j := i + patLen

			for j+patLen <= n && line[j:j+patLen] == pattern {
				count++
				j += patLen
			}

			if count > 1 {
				result.WriteString(fmt.Sprintf("[%d %s]", count, pattern))
				i += count * patLen
				found = true
				break
			}
		}

		if !found {
			result.WriteByte(line[i])
			i++
		}
	}

	return result.String()
}

func EncodeSingleLine(input string) string {
	return encodeLine(input)
}

func EncodeMultiLine(input string) string {
	lines := strings.Split(input, "\n")
	var result []string

	for _, line := range lines {
		result = append(result, encodeLine(line))
	}

	return strings.Join(result, "\n")
}
