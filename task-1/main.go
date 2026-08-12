package main

import (
	"crypto/sha256"
	"fmt"
)

func ttos(value any) string {
	return fmt.Sprintf("%T", value)
}

func vtoa(value any) string {
	return fmt.Sprintf("%v", value)
}

func combine(values ...any) string {
	result := ""
	for _, value := range values {
		result += vtoa(value)
	}
	return result
}

func toRunes(str string) []rune {
	return []rune(str)
}

func addSalt(runes []rune, salt string) []rune {
	return append(runes[:len(runes)/2], append([]rune(salt), runes[len(runes)/2:]...)...)
}

func main() {

	var numDecimal int = 42           // Десятичная система
	var numOctal int = 052            // Восьмеричная система
	var numHexadecimal int = 0x2A     // Шестнадцатиричная система
	var pi float64 = 3.14             // Тип float64
	var name string = "Golang"        // Тип string
	var isActive bool = true          // Тип bool
	var complexNum complex64 = 1 + 2i // Тип complex64

	fmt.Printf("numDecimal type is %s \n", ttos(numDecimal))
	fmt.Printf("numOctal type is %s \n", ttos(numOctal))
	fmt.Printf("numHexadecimal type is %s \n", ttos(numHexadecimal))
	fmt.Printf("pi type is %s \n", ttos(pi))
	fmt.Printf("name type is %s \n", ttos(name))
	fmt.Printf("isActive type is %s \n", ttos(isActive))
	fmt.Printf("complexNum type is %s \n", ttos(complexNum))

	fmt.Printf("Combined: %s \n", combine(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum))
	fmt.Println(toRunes(combine(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)))
	fmt.Println(sha256.Sum256([]byte(string(addSalt(toRunes(combine(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)), "")))))
}
