package main

import (
	"fmt"
	"os"
)

func main() {
	var year int
	fmt.Print("Введите год: ")
	fmt.Fscan(os.Stdin, &year)

	if (year%400 == 0) || (year%4 != 0) {
		fmt.Println("Не високосный")
	} else {
		fmt.Println("Високосный")
	}
}
