package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Print("Введите кол-во строк: ")
	fmt.Fscan(os.Stdin, &n)
	fmt.Println(1)
	tr(1, n, []int{1})

}

func tr(l, n int, prline []int) {
	if l < n {

		var line []int = []int{prline[0]}

		for j := 0; j < l-1; j++ {
			line = append(line, prline[j]+prline[j+1])
		}
		line = append(line, prline[l-1])
		for _, x := range line {
			fmt.Print(x, " ")
		}
		fmt.Println("")
		tr(l+1, n, line)
	}
}
