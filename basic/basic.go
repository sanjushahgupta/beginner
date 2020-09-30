package basic

import "fmt"

var sum, a, b int

func Cal(number1, number2 int) int {
	var addition, subtraction, multiplication, Mod, Division int
	addition = number1 + number2

	subtraction = number1 - number2

	multiplication = number1 * number2

	Mod = number1 / number2

	Division = number1 % number2

	fmt.Printf("%d + %d = %d", number1, number2, addition)
	fmt.Println()
	fmt.Printf("%d - %d = %d", number1, number2, subtraction)
	fmt.Println()
	fmt.Printf("%d * %d = %d", number1, number2, multiplication)
	fmt.Println()
	fmt.Printf("%d modulos %d  = %d", number1, number2, Mod)
	fmt.Println()
	fmt.Printf("%d / %d = %d", number1, number2, Division)
	fmt.Println()
	return 0
}

//func add(a, b int) int {
//sum = a + b
//fmt.Printf("%d + %d = %d", a, b, sum)
//return 0
//}
