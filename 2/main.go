package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

// Приоритет операций
var precedence = map[rune]int{
	'+': 1,
	'-': 1,
	'*': 2,
	'/': 2,
}

// Преобразование инфиксного выражения в ОПЗ
func toRPN(expression string) ([]string, error) {
	var output []string
	var operators []rune

	for i := 0; i < len(expression); i++ {
		c := rune(expression[i])

		// Если число, добавляем в выход
		if unicode.IsDigit(c) || c == '.' {
			numStr := string(c)
			for i+1 < len(expression) && (unicode.IsDigit(rune(expression[i+1])) || expression[i+1] == '.') {
				i++
				numStr += string(expression[i])
			}
			output = append(output, numStr)
		} else if c == '(' {
			// Открывающая скобка в стек
			operators = append(operators, c)
		} else if c == ')' {
			// Закрывающая скобка: извлекаем операции до открывающей
			for len(operators) > 0 && operators[len(operators)-1] != '(' {
				output = append(output, string(operators[len(operators)-1]))
				operators = operators[:len(operators)-1]
			}
			// Убираем открывающую скобку
			if len(operators) == 0 || operators[len(operators)-1] != '(' {
				return nil, fmt.Errorf("несоответствие скобок")
			}
			operators = operators[:len(operators)-1]
		} else if precedence[c] > 0 {
			// Если оператор, учитываем приоритет
			for len(operators) > 0 && precedence[operators[len(operators)-1]] >= precedence[c] {
				output = append(output, string(operators[len(operators)-1]))
				operators = operators[:len(operators)-1]
			}
			operators = append(operators, c)
		} else if c != ' ' {
			// Если неизвестный символ
			return nil, fmt.Errorf("неизвестный символ: %c", c)
		}
	}

	// Добавляем оставшиеся операторы
	for len(operators) > 0 {
		if operators[len(operators)-1] == '(' {
			return nil, fmt.Errorf("несоответствие скобок")
		}
		output = append(output, string(operators[len(operators)-1]))
		operators = operators[:len(operators)-1]
	}

	return output, nil
}

// Вычисление выражения в ОПЗ
func evaluateRPN(rpn []string) (float64, error) {
	var stack []float64

	for _, token := range rpn {
		if num, err := strconv.ParseFloat(token, 64); err == nil {
			// Если число, добавляем в стек
			stack = append(stack, num)
		} else {
			// Если оператор, извлекаем два последних числа из стека
			if len(stack) < 2 {
				return 0, fmt.Errorf("недостаточно операндов для операции %s", token)
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var result float64
			switch token {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				if b == 0 {
					return 0, fmt.Errorf("деление на ноль")
				}
				result = a / b
			default:
				return 0, fmt.Errorf("неизвестный оператор: %s", token)
			}

			// Результат возвращаем в стек
			stack = append(stack, result)
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("ошибка вычисления")
	}

	return stack[0], nil
}

func main() {
	// Ввод выражения
	fmt.Println("Введите арифметическое выражение:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	expression := scanner.Text()

	// Преобразование в ОПЗ
	rpn, err := toRPN(expression)
	if err != nil {
		fmt.Println("Ошибка преобразования в ОПЗ:", err)
		return
	}

	// Вычисление результата
	result, err := evaluateRPN(rpn)
	if err != nil {
		fmt.Println("Ошибка вычисления:", err)
		return
	}

	fmt.Printf("Результат: %.2f\n", result)
}
