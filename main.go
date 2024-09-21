package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	var allRomanArabic = map[string]int{
		"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5,
		"6": 6, "7": 7, "8": 8, "9": 9, "10": 10,
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
		"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}
	/* у меня римские и арабские цифры были записаны в тип string по этому я "1" = 1 тоже конвертирую
	возможно это можно было сделать как-то проще, но я понял как это сделать таким образом
	*/

	var convArabicToRoman = map[int]string{
		1: "I", 2: "II", 3: "III", 4: "IV",
		5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX",
		10: "X", 11: "XI", 12: "XII", 13: "XIII", 14: "XIV",
		15: "XV", 16: "XVI", 17: "XVII", 18: "XVIII", 19: "XIX",
		20: "XX", 21: "XXI", 22: "XXII", 23: "XXIII", 24: "XXIV",
		25: "XXV", 26: "XXVI", 27: "XXVII", 28: "XXVIII", 29: "XXIX",
		30: "XXX", 31: "XXXI", 32: "XXXII", 33: "XXXIII", 34: "XXXIV",
		35: "XXXV", 36: "XXXVI", 37: "XXXVII", 38: "XXXVIII", 39: "XXXIX",
		40: "XL", 41: "XLI", 42: "XLII", 43: "XLIII", 44: "XLIV",
		45: "XLV", 46: "XLVI", 47: "XLVII", 48: "XLVIII", 49: "XLIX",
		50: "L", 51: "LI", 52: "LII", 53: "LIII", 54: "LIV",
		55: "LV", 56: "LVI", 57: "LVII", 58: "LVIII", 59: "LIX",
		60: "LX", 61: "LXI", 62: "LXII", 63: "LXIII", 64: "LXIV",
		65: "LXV", 66: "LXVI", 67: "LXVII", 68: "LXVIII", 69: "LXIX",
		70: "LXX", 71: "LXXI", 72: "LXXII", 73: "LXXIII", 74: "LXXIV",
		75: "LXXV", 76: "LXXVI", 77: "LXXVII", 78: "LXXVIII", 79: "LXXIX",
		80: "LXXX", 81: "LXXXI", 82: "LXXXII", 83: "LXXXIII", 84: "LXXXIV",
		85: "LXXXV", 86: "LXXXVI", 87: "LXXXVII", 88: "LXXXVIII", 89: "LXXXIX",
		90: "XC", 91: "XCI", 92: "XCII", 93: "XCIII", 94: "XCIV",
		95: "XCV", 96: "XCVI", 97: "XCVII", 98: "XCVIII", 99: "XCIX",
		100: "C",
	}
	println("Добро пожаловать в калькулятор")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// создаём регулярное выражение для поиска чисел и операторов
	reAll := regexp.MustCompile(`\d+|[IIIVXCLDM]+|\+|\-|\*|\/|\/`)
	// поиск всех совпадений по инструкции выше
	slice := reAll.FindAllString(input, -1)

	if len(slice) != 3 {
		panic("Ошибка неправильный ввод, использ больше 2-х чисел, \n попытка получения отрицательного числа в римской системе ")
		// Исправил ошибку меня переклинило и я неправильный символ ставил...
	}
	theArabic := regexp.MustCompile(`\d+`)         // проверка на араб числа
	theRoman := regexp.MustCompile(`[IIIVXCLDM]+`) // проверка на рим числа

	reArabic := true // изначально Арабское число равно истине
	reRoman := true  // изначально римское число равно истине

	// цикл фор мы берём слайс и проходим по каждому элементу
	for _, elem := range slice {
		if theArabic.MatchString(elem) { // проверяем является ли наши числа только арабскими
			reRoman = false // если числа только арабские то выставляем флажок на римские = False
		} else if theRoman.MatchString(elem) { // всё тоже самое, но меняем местами сначал рим и проверяем нет ли в них арабских
			reArabic = false //
		} else if elem == "+" || elem == "-" || elem == "*" || elem == "/" { // проверка операторов
			continue
		} else if elem != "0" {
			panic("одно или несколько числе равны 0")
		} else {
			panic("Ошибка: неверный оператор или смешение типов")
			return
		}
	}

	if reArabic && !reRoman { // тут мы уже сравниваем наши булевы значения один должен быть только False а другой только True
		// fmt.Println("Input содержит только арабские числа")
	} else if reRoman && !reArabic {
		// fmt.Println("Input содержит только римские числа")
	} else { // но а если оба варианта True то у нас было смешивание типов чисел
		panic("Ошибка: смешивание арабских и римских чисел")
	}

	num1 := slice[0]
	oper := slice[1]
	num2 := slice[2]
	val1, ok1 := allRomanArabic[num1]
	val2, ok2 := allRomanArabic[num2]

	if val1 == 0 {
		panic(" 0")
		return
	}
	if val2 == 0 {
		panic(" 0")
		return
	}

	if !ok1 || !ok2 {
		panic("Одно из чисел некорректное или вы пытаетесь поделить на 0")
		return
	}

	var result int
	switch oper {
	case "+":
		result = val1 + val2
	case "-":

		result = val1 - val2

	case "*":

		result = val1 * val2
	case "/":
		if val2 == 0 {
			panic("Ошибка деления на 0")
		}
		result = val1 / val2

	}
	if reRoman && !reArabic { // обрабатываем римские числа
		if result < 1 {
			panic("Ошибка: результат меньше или равен 0")
		}
		if romanResult, exists := convArabicToRoman[result]; exists {
			fmt.Println(romanResult)
		}
	} else { // обрабатываем арабские числа
		fmt.Println(result)
	}

}
