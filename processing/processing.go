package processing

import (
	"sync"
)

// ProcessData processes each line in parallel: decodes or encodes.
func ProcessData(lines []string, encode bool) []string {
	result := make([]string, len(lines))
	var wg sync.WaitGroup
	wg.Add(len(lines))

	for i, line := range lines {
		go func(i int, l string) {
			defer wg.Done()
			if encode {
				result[i] = Encode(l)
			} else {
				decoded, err := Decode(l)
				if err != nil {
					result[i] = "Error"
				} else {
					result[i] = decoded
				}
			}
		}(i, line)
	}
	wg.Wait()
	return result
}
