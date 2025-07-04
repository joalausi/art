package processing

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// either literal or [N STR]
type Token struct {
	Repeat int
	Text   string
}

// Parses a string into tokens or returns an error
func Parse(input string) ([]Token, error) {
	// regexp for matching [N STR] patterns
	// where N is a number and STR is any non-empty string
	re := regexp.MustCompile(`\[(\d+)\s([^\]]+)\]`)
	tokens := []Token{}
	lastIndex := 0

	for _, match := range re.FindAllStringSubmatchIndex(input, -1) {
		start, end := match[0], match[1]
		// add everything before [
		if start > lastIndex {
			tokens = append(tokens, Token{Repeat: 1, Text: input[lastIndex:start]})
		}
		// parse number and string
		num, err := strconv.Atoi(input[match[2]:match[3]])
		if err != nil {
			return nil, fmt.Errorf("not a number: %w", err)
		}
		str := input[match[4]:match[5]]
		if len(str) == 0 {
			return nil, fmt.Errorf("empty line in brackets")
		}
		tokens = append(tokens, Token{Repeat: num, Text: str})
		lastIndex = end
	}

	if lastIndex < len(input) {
		tokens = append(tokens, Token{Repeat: 1, Text: input[lastIndex:]})
	}
	// check for single parentheses and overall balance
	for _, t := range tokens {
		if t.Repeat == 1 && (strings.Contains(t.Text, "[") || strings.Contains(t.Text, "]")) {
			return nil, fmt.Errorf("unbalanced brackets in token: %s", t.Text)
		}
	}
	if strings.Count(input, "[") != strings.Count(input, "]") {
		return nil, fmt.Errorf("unbalanced brackets in input")
	}

	return tokens, nil
}

// ParseMultiLine splits the input into lines and parses each one separately.
func ParseMultiLine(input string) ([][]Token, error) {
	lines := strings.Split(strings.TrimRight(input, "\n"), "\n")
	result := make([][]Token, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			result = append(result, []Token{})
			continue
		}
		tks, err := Parse(line)
		if err != nil {
			return nil, err
		}
		result = append(result, tks)
	}
	return result, nil
}
