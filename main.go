package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Введите предложение")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	text = strings.TrimSpace(text)
	frequency := countCharFrequency(text)
	fmt.Println("Частота символов:")
	for char, count := range frequency {
		fmt.Printf("%c - %.0f %.2f\n", char, count[0], count[1])
	}
}

func countCharFrequency(str string) map[rune][]float32 {
	frequency := make(map[rune][]float32)
	totalChars := float32(0)
	for _, char := range str {
		if unicode.IsLetter(char) {
			lowChar := unicode.ToLower(char)
			if len(frequency[lowChar]) < 1 {
				frequency[lowChar] = []float32{0, 0}
			}
			frequency[lowChar][0]++
			totalChars++
		}
	}
	for _, count := range frequency {
		count[1] = count[0] / totalChars
	}
	return frequency
}
