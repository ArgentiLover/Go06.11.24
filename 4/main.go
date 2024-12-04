package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите строку с пробелами:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	words := strings.Split(input, " ")

	min_size := 1000000
	for _, x := range words {
		min_size = min(min_size, utf8.RuneCountInString(x))
	}
	max_pref := 0
	for i := 0; i < min_size; i++ {
		cur_char := words[0][i]
		isSame := true
		for j := 1; j < len(words); j++ {
			if words[j][i] != cur_char {
				print("Result: ", words[0][:max_pref])
				isSame = false
			}
		}
		if isSame {
			max_pref = max_pref + 1
			if max_pref == min_size {
				print("Result: ", words[0][:max_pref])
			}
		} else {
			break
		}
	}
}
