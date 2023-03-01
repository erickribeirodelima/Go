package main

import "fmt"

func main() {
	var variavel1 string = "VAR1"
	variavel2 := "VAR2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "VAR3"
		variavel4 string = "VAR4"
	)

	fmt.Println(variavel3)
	fmt.Println(variavel4)

	variavel5, variavel6 := "VAR5", "VAR6"
	fmt.Println(variavel5, variavel6)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)
}
