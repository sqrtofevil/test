package main

import (
	"errors" //для возвращения сообщения об ошибке
	"fmt"
)

func count(a, b float64, c interface{}) (float64, error) { //фукнкция проверки операции
	var err error
	var res float64
	switch c.(string) {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	default:
		err = errors.New("неизвестная операция") //если любое другое значение=присваиваем код ошибки
	}
	return res, err //всегда возвращаем результат и код ошибки (если ошибки нет, то код =0)
}

func CheckFloat(a, b interface{}) bool {
	if _, ok := a.(float64); !ok {
		fmt.Printf("value=%v: %T", a, a)
		return false
	} else if _, ok1 := b.(float64); !ok1 {
		fmt.Printf("value=%v: %T", b, b)
		return false
	}
	return true
}

func readTask() (a, b, c interface{}) {
	return 5.0, 2.0, "*" //тут играемся с параметрами ввода через тип переменной интерфейс
}

func main() {
	value1, value2, operation := readTask() //инициализируем ввод данных в переменные
	if CheckFloat(value1, value2) {
		result, err := count(value1.(float64), value2.(float64), operation) //значение в result, ошибку в err
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%.4f", result)
		}
	}
}
