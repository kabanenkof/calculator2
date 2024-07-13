package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const (
	welcomeMessage = "\nДобро пожаловать в Ката калькулятор!\n" +
		"Данный калькулятор позволяет выполнять простейшие арифметические операции\n" +
		"над арабскими или над римскими цифрами.\n\n" +
		"==========================================================================\n" +
		"Введиите данные для рассчета в следующем формате: Число (арабское/римское) | Знак математической операции (+, -, *, /) | Число (арабское/римское)\n"

	typeErrorMessage                   = "Введите ОДИНАКОВЫЕ типы цифр(арабские/римские)."
	incorrectInputErrorMessage         = "Введено что-то некорректное (тип ввода: число, знак мат операции, число)."
	incorrectValueErrorMessage         = "Необходимо использовать числа от 1 до 10."
	incorrectMathOperationErrorMessage = "Введите верный знак математической операции (+, -, *, /)."
	incorrectResultErrorMessage        = "Резуультат меньше 1."
)

func main() {
	var (
		firstNumber, secondNumber, calculateResult            int
		firstStrNumber, secondStrNumber, mathOperationStrSign string
		err                                                   = ""
	)

	fmt.Println(welcomeMessage)
	fmt.Fscanln(os.Stdin, &firstStrNumber, &mathOperationStrSign, &secondStrNumber, &err)

	if err != "" {
		panic(incorrectInputErrorMessage)
		return
	}

	if inputIsRoman(firstStrNumber) != inputIsRoman(secondStrNumber) {
		panic(typeErrorMessage)
		return
	}

	//Convert strNumbers -> Int
	if inputIsRoman(firstStrNumber) {
		firstNumber = romanToArabic(firstStrNumber)
	} else {
		firstNumber, _ = strconv.Atoi(firstStrNumber)
	}
	if inputIsRoman(secondStrNumber) {
		secondNumber = romanToArabic(secondStrNumber)
	} else {
		secondNumber, _ = strconv.Atoi(secondStrNumber)
	}

	//Check numbers is > 0
	if secondNumber > 10 || secondNumber < 1 || firstNumber > 10 || firstNumber < 1 {
		panic(incorrectValueErrorMessage)
		return
	}

	switch mathOperationStrSign {
	case "+":
		calculateResult = firstNumber + secondNumber
	case "-":
		calculateResult = firstNumber - secondNumber
	case "*":
		calculateResult = firstNumber * secondNumber
	case "/":
		calculateResult = firstNumber / secondNumber
	default:
		panic(incorrectMathOperationErrorMessage)
		return
	}

	if inputIsRoman(firstStrNumber) && inputIsRoman(secondStrNumber) {
		var chekResult = arabicToRoman(calculateResult)
		if chekResult == "" {
			panic(incorrectResultErrorMessage)
			return
		}
		fmt.Println("Результат вычисления: ", arabicToRoman(calculateResult))
	} else {
		fmt.Println("Результат вычисления: ", calculateResult)
	}

}

// Convert roman -> arabic
func romanToArabic(romanStrNumber string) int {
	var romanNumbers = map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7,
		"VIII": 8, "IX": 9, "X": 10, "L": 50, "C": 100,
	}

	return romanNumbers[romanStrNumber]
}

// Convert arabic -> roman
func arabicToRoman(arabicNumber int) string {
	var convertResult strings.Builder

	romanSymbols := []struct {
		value  int
		symbol string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	for _, symbol := range romanSymbols {
		for arabicNumber >= symbol.value {
			convertResult.WriteString(symbol.symbol)
			arabicNumber -= symbol.value
		}
	}

	return convertResult.String()
}

// Check input string for valid
func inputIsRoman(inputNumber string) bool {
	for _, inputChar := range inputNumber {
		if !unicode.IsLetter(inputChar) || (unicode.ToUpper(inputChar) != 'I' &&
			unicode.ToUpper(inputChar) != 'V' &&
			unicode.ToUpper(inputChar) != 'X') {
			return false
		}
	}
	return true
}
