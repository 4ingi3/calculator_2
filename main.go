package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Ввод всего выражения
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите выражение (например, \"i\" * 3): ")
	expression, _ := reader.ReadString('\n')
	expression = strings.TrimSpace(expression)

	// Убираем пробелы
	expression = strings.ReplaceAll(expression, " ", "")

	// Определяем оператор и разбиваем строку на части
	var operator string
	for _, op := range []string{"*", "+", "-", "/"} {
		if strings.Contains(expression, op) {
			operator = op
			break
		}
	}

	if operator == "" {
		fmt.Println("Ошибка: неподдерживаемый оператор.")
		return
	}

	parts := strings.Split(expression, operator)
	if len(parts) != 2 {
		fmt.Println("Ошибка: выражение должно состоять из двух частей.")
		return
	}

	// Проверка длины первой строки
	str1 := getStringArg(parts[0])
	if str1 == "" {
		fmt.Println("Ошибка: первый аргумент должен быть строкой в кавычках.")
		return
	}

	if len(str1) > 10 {
		fmt.Println("Ошибка: строка не должна содержать более 10 символов.")
		return
	}

	str2 := parts[1]
	var result string

	switch operator {
	case "+":
		str2 = getStringArg(str2)
		if str2 == "" {
			fmt.Println("Ошибка: второй аргумент должен быть строкой в кавычках.")
			return
		}
		if len(str2) > 10 {
			fmt.Println("Ошибка: строка не должна содержать более 10 символов.")
			return
		}
		result = str1 + str2
	case "-":
		str2 = getStringArg(str2)
		if str2 == "" {
			fmt.Println("Ошибка: второй аргумент должен быть строкой в кавычках.")
			return
		}
		result = removeSubstring(str1, str2)
	case "*":
		num, valid := stringToInt(str2)
		if !valid || num < 1 || num > 10 {
			fmt.Println("Ошибка: второй аргумент должен быть числом от 1 до 10.")
			return
		}
		result = repeatString(str1, num)
	case "/":
		num, valid := stringToInt(str2)
		if !valid || num < 1 || num > 10 {
			fmt.Println("Ошибка: второй аргумент должен быть числом от 1 до 10.")
			return
		}
		result = divideString(str1, num)
	}

	if len(result) > 40 {
		result = result[:40] + "..."
	}

	fmt.Println(result)
}

// Функция для получения строки без кавычек
func getStringArg(s string) string {
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		return s[1 : len(s)-1]
	}
	return ""
}

// Преобразуем строку в целое число
func stringToInt(s string) (int, bool) {
	result := 0
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		result = result*10 + int(s[i]-'0')
	}
	return result, true
}

// Повторяем строку n раз
func repeatString(s string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += s
	}
	return result
}

// Делим строку на n частей
func divideString(s string, n int) string {
	if n > 0 && len(s) >= n {
		return s[:len(s)/n]
	}
	return ""
}

// Удаляем подстроку из строки
func removeSubstring(s, sub string) string {
	return strings.Replace(s, sub, "", -1)
}
