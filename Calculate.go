package Calculator

import (
	"fmt"
)

//This function receives the arithmetic operation symbol
func GetArithmeticSymbol(symbol string) string {
	fmt.Println("Введите '+', либо, '-' , либо '*', либо '/' ")

	_, err := fmt.Scan(&symbol)

	if err != nil {
		fmt.Println("Неправильный формат ввода!")
	} else if symbol != "+" && symbol != "-" && symbol != "*" && symbol != "/" {
		fmt.Println("Неправильный формат ввода!")
		return GetArithmeticSymbol(symbol)
	}
	return symbol
}

func AddOrSubtract() float64 {
	var (
		v1, v2, result float64
		symbol         string
	)

	GetNumber1(&v1)

	symbol = GetArithmeticSymbol(symbol)

	GetNumber2(&v2)

	switch symbol {
	case "+":
		result = v1 + v2
	case "-":
		result = v1 - v2
	case "*":
		result = v1 * v2
	case "/":
		if v2 == 0 {
			fmt.Println("Деление на 0 запрещено !!!")
			GetNumber2(&v2)
			result = v1 / v2
		} else {
			result = v1 / v2
		}
		fmt.Println(v1, symbol, v2, "=", result)
	}

	return result
}
