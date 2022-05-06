package Calculator

import "fmt"

//This function receives the first number
func GetNumber1(number *float64) float64 {

	fmt.Println("Введите первое число")

	_, err := fmt.Scan(number)
	if err != nil {
		fmt.Println("Неправильный формат ввода! Попробуйте еще раз!")
		return GetNumber1(number)
	}

	return *number
}

//This function receives the second number
func GetNumber2(number *float64) float64 {

	fmt.Println("Введите второе число")

	_, err := fmt.Scan(number)
	if err != nil {
		fmt.Println("Неправильный формат ввода! Попробуйте еще раз!")
		return GetNumber2(number)
	}

	return *number
}
