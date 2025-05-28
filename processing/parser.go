package parser

import (
    "fmt"
    "regexp"
    "strconv"
)

type Token struct {
    Repeat int
    Text   string
}
func Parse(input string) ([]Token, error) {
    re := regexp.MustCompile(`\[(\d+)\s([^\]]+)\]`)
    tokens := []Token{}
    lastIndex := 0

    for _, match := range re.FindAllStringSubmatchIndex(input, -1) {
        start, end := match[0], match[1]
        if start > lastIndex {
            tokens = append(tokens, Token{Repeat: 1, Text: input[lastIndex:start]})
        }
        num, err := strconv.Atoi(input[match[2]:match[3]])
        if err != nil {
            return nil, fmt.Errorf("не число: %w", err)
        }
        str := input[match[4]:match[5]]
        if len(str) == 0 {
            return nil, fmt.Errorf("пустая строка в скобках")
        }
        tokens = append(tokens, Token{Repeat: num, Text: str})
        lastIndex = end
    }
    if lastIndex < len(input) {
        tokens = append(tokens, Token{Repeat: 1, Text: input[lastIndex:]})
    }
    if countOpen := len(re.FindAllString(input, -1)); countOpen*2 != len(re.FindAllString(input, ""]")) {
//         // тут упрощённо; можно более строго
//         // но regexp не найдёт несбалансированные, если лишняя ]
//         // так что можно досканировать на наличие одиночных ]
//         if regexp.MustCompile(`\]|\[`).MatchString(input) {
//             return nil, fmt.Errorf("несбалансированные скобки")
//         }
//     }

//     return tokens, nil
// }
