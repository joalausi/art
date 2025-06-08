package processing

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var pattern = regexp.MustCompile(`\[(\d+) (.+?)\]`)

func DecodeSingleLine(input string) (string, error) {
	result := ""

	for {
		match := pattern.FindStringSubmatchIndex(input)
		if match == nil {
			result += input
			break
		}

		start, end := match[0], match[1]
		countStr := input[match[2]:match[3]]
		val := input[match[4]:match[5]]

		count, err := strconv.Atoi(countStr)
		if err != nil || val == "" {
			return "", errors.New("invalid input")
		}

		result += input[:start]
		result += strings.Repeat(val, count)
		input = input[end:]
	}

	if strings.ContainsAny(result, "[]") {
		return "", errors.New("unbalanced brackets")
	}

	return result, nil
}

func DecodeMultiLine(input string) (string, error) {

	lines := strings.Split(input, "\n")
	var result []string

	for _, line := range lines {
		decodedline, err := DecodeSingleLine(line)
		if err != nil {
			return "", err
		}
		result = append(result, decodedline)
	}
	return strings.Join(result, "\n"), nil
}

// Decode is a thin wrapper around DecodeSingleLine. It exists so the
// ProcessData helper can use a common function name for both encoding and
// decoding. The current CLI processes one line at a time, so this helper simply
// delegates to DecodeSingleLine.
func Decode(input string) (string, error) {
	return DecodeSingleLine(input)
}
