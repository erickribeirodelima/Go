package main

import "fmt"

func main() {
	// Aritméticos
	soma := 2 + 2
	subtracao := 6 - 2
	multiplicacao := 2 * 2
	divisao := 8 / 2
	restoDaDivisao := 10 % 2

	fmt.Println(soma)
	fmt.Println(subtracao)
	fmt.Println(multiplicacao)
	fmt.Println(divisao)
	fmt.Println(restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 10
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	// Fim dos aritméticos

	// Atribuição
	var nome string = "Erick"
	sobrenome := "Ribeiro"
	fmt.Println(nome, sobrenome)
	// Fim operadores de atribuição

	// Operadores Relacionais
	// Obs: Operadores Relacionais sempre retornam um booleano

	fmt.Println(1 > 2)  // false
	fmt.Println(2 < 1)  // false
	fmt.Println(2 >= 1) // true
	fmt.Println(1 <= 2) // true
	fmt.Println(2 == 2) // true
	fmt.Println(2 != 1) // true

	// Fim dos relacionais

	// Operadores lógicos
	// apenas 3

	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso) // false
	fmt.Println(verdadeiro || falso) // true
	fmt.Println(!verdadeiro)         // false

	// Fim operadores lógicos

	// Operadores Unários
	// incrementar
	numero := 10
	numero++
	fmt.Println(numero) // 11
	numero += 15
	fmt.Println(numero) // 26
	// decrementar
	numero--
	fmt.Println(numero) // 25
	// Fim dos operadores unários

}
