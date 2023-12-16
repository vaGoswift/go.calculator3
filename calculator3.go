package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите выражение: ")
	expression, _ := reader.ReadString('\n')

	result, err := calculate(expression)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("Результат:", result)
}

func calculate(expression string) (string, error) {
	tokens := strings.Fields(expression)
	if len(tokens) != 3 {
		return "", fmt.Errorf("формат математической операции должен быть: число оператор число")
	}

	a, err := parseNumber(tokens[0])
	if err != nil {
		return "", err
	}

	operator := tokens[1]

	b, err := parseNumber(tokens[2])
	if err != nil {
		return "", err
	}

	isRomanA := isRomanNumeral(tokens[0])
	isRomanB := isRomanNumeral(tokens[2])
	if isRomanA != isRomanB {
		return "", fmt.Errorf("используются разные системы счисления")
	}

	result := 0
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			return "", fmt.Errorf("деление на ноль")
		}
		result = a / b
	default:
		return "", fmt.Errorf("неверный оператор, используйте: +, -, *, /")
	}

	if isRomanA {
		if result <= 0 {
			return "", fmt.Errorf("результат работы с римскими числами должен быть положительным")
		}
		return toRoman(result), nil
	}

	return strconv.Itoa(result), nil
}

func parseNumber(str string) (int, error) {
	if val, err := strconv.Atoi(str); err == nil {
		if val < 1 || val > 10 {
			return 0, fmt.Errorf("число должно быть в диапазоне от 1 до 10")
		}
		return val, nil
	}

	romanNumerals := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
		"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}

	val, ok := romanNumerals[str]
	if !ok {
		return 0, fmt.Errorf("неверное число")
	}

	return val, nil
}

func isRomanNumeral(str string) bool {
	romanNumerals := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}

	for _, numeral := range romanNumerals {
		if str == numeral {
			return true
		}
	}

	return false
}

func toRoman(n int) string {
	romanValues := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanNumerals := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var result strings.Builder

	for i := 0; i < len(romanValues); i++ {
		for n >= romanValues[i] {
			result.WriteString(romanNumerals[i])
			n -= romanValues[i]
		}
	}

	return result.String()
}
