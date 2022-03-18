package Calculator

import "fmt"

func plusOrMinus(symbol string) string {

	fmt.Println("Введите +/-")

	_, err := fmt.Scan(&symbol)

	if err != nil {
		fmt.Println("Неправильный формат ввода!")
	} else if symbol != "+" && symbol != "-" {
		fmt.Println("Неправильный формат ввода!")
		return plusOrMinus(symbol)
	}

	return symbol
}

func AddOrSubtract() float64{
	var (
		v1, v2, result float64
		symbol string
	)

	//Этот кусок кода принимает первое значение
	fmt.Println("Введите первое число")
	_, err := fmt.Scan(&v1)
	if err != nil {
		fmt.Println("Неправильный формат ввода!")
		return AddOrSubtract()
	}

	//Вводите или + или -
	symbol = plusOrMinus(symbol)

	//Этот кусок кода принимает первое значение
	fmt.Println("Введите второе число")
	_, err = fmt.Scan(&v2)
	if err != nil {
		fmt.Println("Неправильный формат ввода!")
		return AddOrSubtract()
	}

	//В зависимости от выбранного символа, два значения или складываются или отнимаются
	if symbol == "+" {
		result = v1 + v2
	}else if symbol == "-"{
		result = v1 - v2
	}

	fmt.Println(v1, symbol, v2, "=", result)

	return result
}



