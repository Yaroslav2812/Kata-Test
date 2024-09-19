package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	allRomanArabic := make(map[string]int)
	allRomanArabic["I"] = 1
	allRomanArabic["II"] = 2
	allRomanArabic["III"] = 3
	allRomanArabic["IV"] = 4
	allRomanArabic["V"] = 5
	allRomanArabic["VI"] = 6
	allRomanArabic["VII"] = 7
	allRomanArabic["VIII"] = 8
	allRomanArabic["IX"] = 9
	allRomanArabic["X"] = 10
	allRomanArabic["1"] = 1
	allRomanArabic["2"] = 2
	allRomanArabic["3"] = 3
	allRomanArabic["4"] = 4
	allRomanArabic["5"] = 5
	allRomanArabic["6"] = 6
	allRomanArabic["7"] = 7
	allRomanArabic["8"] = 8
	allRomanArabic["9"] = 9
	allRomanArabic["10"] = 10
	/* у меня римские и арабские цифры были записаны в тип string по этому я "1" = 1 тоже конвертирую
	возможно это можно было сделать как-то проще, но я понял как это сделать таким образом
	*/

	convArabicToRoman := make(map[int]string)
	convArabicToRoman[1] = "I"
	convArabicToRoman[2] = "II"
	convArabicToRoman[3] = "III"
	convArabicToRoman[4] = "IV"
	convArabicToRoman[5] = "V"
	convArabicToRoman[6] = "VI"
	convArabicToRoman[7] = "VII"
	convArabicToRoman[8] = "VIII"
	convArabicToRoman[9] = "IX"
	convArabicToRoman[10] = "X"
	convArabicToRoman[11] = "XI"
	convArabicToRoman[12] = "XII"
	convArabicToRoman[13] = "XIII"
	convArabicToRoman[14] = "XIV"
	convArabicToRoman[15] = "XV"
	convArabicToRoman[16] = "XVI"
	convArabicToRoman[17] = "XVII"
	convArabicToRoman[18] = "XVIII"
	convArabicToRoman[19] = "XIX"
	convArabicToRoman[20] = "XX"

	println("Добро пожаловать в калькулятор")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// создаём регулярное выражение для поиска чисел и операторов
	reAll := regexp.MustCompile(`\d+|[IIIVXCLDM]+|\+|\-|\*|\/`)
	// поиск всех совпадений по инструкции выше
	slice := reAll.FindAllString(input, -1)

	if len(slice) != 3 {
		fmt.Println("Ошибка неправильный ввод 3 + 3 + 3. правильный ввод 3 + 3  ")
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
		} else {
			println("Ошибка неверный оператор или смешнные числа")
			return
		}
	}

	if reArabic && !reRoman { // тут мы уже сравниваем наши булевы значения один должен быть только False а другой только True
		// fmt.Println("Input содержит только арабские числа")
	} else if reRoman && !reArabic {
		// fmt.Println("Input содержит только римские числа")
	} else { // но а если оба варианта True то у нас было смешивание типов чисел
		fmt.Println("Ошибка: смешивание арабских и римских чисел")

	}

	num1 := slice[0]
	oper := slice[1]
	num2 := slice[2]
	val1, ok1 := allRomanArabic[num1]
	val2, ok2 := allRomanArabic[num2]

	if !ok1 || !ok2 {
		println("Одно из чисел некорректное или вы пытаетесь поделить на 0")

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
		result = val1 / val2
	}

	reeRoman := regexp.MustCompile(`[IIIVXCLDM]+`)
	reeArabic := regexp.MustCompile(`\d+`)
	if reeRoman.MatchString(input) {
		resultConv := result
		value, exists := convArabicToRoman[resultConv]
		if reRoman && !reArabic {

			fmt.Println(value)
		} else if exists {

		}

	} else if reeArabic.MatchString(input) {
		println(result)
	}

	// последний цикл я сам не понял как он работает, но он работает
	// мы в цикле проверяем какие у нас цифры римские или арабские если арабские то просто выводим ответ
	// если римские то мы через мапу конвертируем его в строку
	// для этого я взял значение из переменной result перезаписал его в новую переменную
	// далее я сделал сверку значений из мапы convArabicToRoman и записал ответ в Value
	// так же была ошибка если я вводил значения 3 + V то он выдавал ошибку и выдавал результат
	// это помогло исправить опять проверка reRoman && !reArabic
	// я думал что это не поможет, но помогло. Дальше я действовал согласно правилу, работает не трогай.
}
