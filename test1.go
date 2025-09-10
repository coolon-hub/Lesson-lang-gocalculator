package main

import "fmt"

var (
	num1      float32
	Sign      string
	num2      float32
	Exit      string
	Answer    float32
	Factorial float32 = 1
	i         float32
)

func main() {

	for {
		menu()
		if Exit == "Нет" || Exit == "нет" {
			break
		}
	}

	fmt.Print("\033[H\033[2J")
	fmt.Println("Всего доброго, Пользователь!")
}

func menu() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Добро пожаловать в калькулятор, Пользователь!!")
	fmt.Print("Введите первое число: ")
	fmt.Scan(&num1)
	fmt.Print("Введите арефметический знак(+,-,*,/):")
	fmt.Scan(&Sign)
	check1()
}

func check1() {
	if Sign == "+" {
		number2()
		Answer = num1 + num2
		fmt.Println("Ответ:", Answer)
	} else if Sign == "-" {
		number2()
		Answer = num1 - num2
		fmt.Println("Ответ:", Answer)
	} else if Sign == "*" {
		number2()
		Answer = num1 * num2
		fmt.Println("Ответ:", Answer)
	} else if Sign == "/" {
		number2()
		Answer = num1 / num2
		fmt.Println("Ответ:", Answer)
	} else {
		check2()
	}
	exitf()
}

func check2() {
	if Sign == "!" {
		for ; i < num1; i++ {
			Factorial *= num1
		}
		fmt.Println("Ответ:", Factorial)
	}
}

func number2() {
	fmt.Print("Введите второе число:")
	fmt.Scan(&num2)
}
func exitf() {
	fmt.Print("Хотите продолжить?:")
	fmt.Scan(&Exit)
}
