package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var num1str, sign, num2str string
	log.Print("Введите задание")
	n, err := fmt.Fscan(os.Stdin, &num1str, &sign, &num2str)
	if err != nil {
		log.Panic("Ошибка чтения")
		//return
	}

	log.Printf("n=%v, err=%v", n, err)

	log.Printf("&task1=%v, &sign=%v, &task2=%v", num1str, sign, num2str)

	//Пытаемся преобразуем в обычные числа
	var num1, num2 int
	var num1IsRoman, num2IsRoman bool

	num1, err = strconv.Atoi(num1str)
	if err != nil {
		log.Printf("Параметр 1 не является числом")
		//return
		num1 = romanToInt(num1str)
		if num1 == 0 {
			log.Panic("Параметр 1 не является римским числом")
			return
		} else {
			log.Printf("Параметр 1 является римским числом %v", num1)
			num1IsRoman = true
		}
	}
	if num1 > 10 {
		log.Panic("параметр один > 10")
		return
	}

	num2, err = strconv.Atoi(num2str)
	if err != nil {
		log.Printf("Параметр 2 не является числом")
		//return
		num2 = romanToInt(num2str)
		if num2 == 0 {
			log.Panic("Параметр 2 не является римским числом")
			return
		} else {
			log.Printf("Параметр 2 является римским числом %v", num2)
			num2IsRoman = true
		}
	}

	if num2 > 10 {
		log.Print("параметр два > 10")
		return
	}

	if (num1IsRoman && !num2IsRoman) || (!num1IsRoman && num2IsRoman) {
		log.Panic("Одно из значений римское, а другое - нет")
		return
	}

	//Используя switch/if-else выполнить операцию указанную во вводе либо выдать ошибку
	var result int
	switch sign {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "/":
		result = num1 / num2
	case "*":
		result = num1 * num2
	default:
		log.Panicf(`Не допустимая операция "%v"`, sign)
		//return
	}

	var resultStr string
	if num1IsRoman && num2IsRoman {
		resultStr = intToRoman(result)
	} else {
		resultStr = fmt.Sprint(result)
	}

	log.Printf("Для задания %v %v %v = %v", num1str, sign, num2str, resultStr)

	return
}

func romanToInt(s string) int {

	romanMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var result, prevValue int

	sr := []rune(s) //Получаем массив рун

	//VI

	for i := len(sr) - 1; i >= 0; i-- {
		currentValue, ok := romanMap[sr[i]]
		if !ok {
			log.Panicf(`недопустимый символ "%v"`, string(sr[i]))
		}
		if currentValue >= prevValue {
			result += currentValue
		} else {
			result -= currentValue
		}

		prevValue = currentValue
	}

	return result
}

func intToRoman(num int) string {
	if num < 1 {
		log.Panic("Римсоке число не может быть < 1")
	}
	// Определение соответствия чисел и символов римской системы
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	// Инициализация переменной для хранения результата
	result := ""

	// Проход по значениям и символам римской системы
	for i := 0; i < len(values); i++ {
		// Пока число больше или равно текущему значению, добавляем символ в результат и вычитаем значение
		for num >= values[i] {
			result += romans[i]
			num -= values[i]
		}
	}

	return result
}
