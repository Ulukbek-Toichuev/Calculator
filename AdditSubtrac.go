package Calculator

import "fmt"

//This function receives the arithmetic operation symbol
func plusOrMinus(symbol string) string {
	fmt.Println("Введите '+', либо, '-' , либо '*', либо '/' ")

	_, err := fmt.Scan(&symbol)

	if err != nil {
		fmt.Println("Неправильный формат ввода!")
	} else if symbol != "+" && symbol != "-" && symbol != "*" && symbol != "/" {
		fmt.Println("Неправильный формат ввода!")
		return plusOrMinus(symbol)
	}
	return symbol
}

//This function receives the first number
func getNumber1(number float64) float64 {
	fmt.Println("Введите первое число")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Неправильный формат ввода! Попробуйте еще раз!")
		return getNumber1(number)
	}
	return number
}

//This function receives the second number
func getNumber2(number float64) float64 {
	fmt.Println("Введите второе число")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Неправильный формат ввода! Попробуйте еще раз!")
		return getNumber2(number)
	}
	return number
}

func AddOrSubtract() float64 {
	var (
		v1, v2, result float64
		symbol         string
	)

	v1 = getNumber1(v1)

	symbol = plusOrMinus(symbol)

	v2 = getNumber2(v2)

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
			return getNumber2(v2)
		}
		result = v1 / v2
	}

	fmt.Println(v1, symbol, v2, "=", result)

	return result
}
