package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func solution(input string, prefix string) string {
	for number := 1; ; number++ {
		content := fmt.Sprintf("%v%v", input, number)
		hash := md5.Sum([]byte(content))
		hashStr := hex.EncodeToString(hash[:])
		if strings.HasPrefix(hashStr, prefix) {
			return fmt.Sprint(number)
		}
	}
}

func runPart1(input string) string {
	return solution(input, "00000")
}

func runPart2(input string) string {
	return solution(input, "000000")
}

func main() {
	input := "iwrupvqb"

	fmt.Println(runPart1(input))
	fmt.Println(runPart2(input))
}
