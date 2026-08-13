package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

func ttos(value any) string {
	return fmt.Sprintf("%T", value)
}

func vtoa(value any) string {
	return fmt.Sprint(value)
}

func combine(values ...any) string {
	var result strings.Builder
	for _, value := range values {
		result.WriteString(vtoa(value))
	}
	return result.String()
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

	fmt.Println("numDecimal type is", ttos(numDecimal))
	fmt.Println("numOctal type is", ttos(numOctal))
	fmt.Println("numHexadecimal type is", ttos(numHexadecimal))
	fmt.Println("pi type is", ttos(pi))
	fmt.Println("name type is", ttos(name))
	fmt.Println("isActive type is", ttos(isActive))
	fmt.Println("complexNum type is", ttos(complexNum))

	fmt.Println("Combined values:", combine(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum))
	fmt.Println("combined in runes:", toRunes(combine(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)))
	fmt.Println("combined runes with salt:", string(addSalt(toRunes(combine(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)), "go-2024")))
	fmt.Println("sha256 sum:", sha256.Sum256([]byte(string(addSalt(toRunes(combine(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)), "go-2024")))))
}
