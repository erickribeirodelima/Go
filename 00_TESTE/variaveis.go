package main

import (
	"fmt"
)

func main() {
	var variavel1 string = "Teste!"
	variavel2 := "Teste 2!"

	fmt.Println(variavel1, variavel2)

	var (
		variavel3 string = "Teste 3"
		variavel4 string = "Teste 4"
	)

	fmt.Println(variavel3)
	fmt.Println(variavel4)

	variavel5, variavel6 := "Teste 5", "Teste 6"

	fmt.Println(variavel5)
	fmt.Println(variavel6)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5

	fmt.Println(variavel5)
	fmt.Println(variavel6)

}
