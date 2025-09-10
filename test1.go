package main

import (
	"fmt"
	"time"
)

var (
	num1      float32
	Sign      string
	num2      float32
	Exit      string
	Answer    float32
	Factorial float32 = 1
	i         float32 = 1
)

func main() {
	fmt.Print("\033[H\033[2J")
	load()
	for {
		menu()
		/* if Exit == "Нет" || Exit == "нет" {
			menu()
		}  else*/if Exit == "Да" || Exit == "да" {
			break
		} /*else {
			fmt.Print("\033[H\033[2J")
			fmt.Println("Вы вели не правильные данные, попробуйте ещё раз.")
			exitf()
		}*/
	}

	fmt.Print("\033[H\033[2J")
	fmt.Println("Всего доброго, Пользователь!")
}

func load() {
	for time1 := 1; time1 <= 30; time1++ {
		for time2 := 1; time2 <= 3; time2++ {
			fmt.Print(".")
			time.Sleep(300 * time.Millisecond)
		}
		fmt.Print("\033[H\033[2J")
	}
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
		for ; i <= num1; i++ {
			Factorial *= i
		}
		fmt.Println("Ответ:", Factorial)
	}
}

func number2() {
	fmt.Print("Введите второе число:")
	fmt.Scan(&num2)
}
func exitf() {
	fmt.Print("Хотите закончить?:")
	fmt.Scan(&Exit)
}
