package main

import (
	"errors"
	"fmt"
)

func main() {
	// -- Números inteiros
	// int8 int16 int32 int64
	// int  (int vazio significa arquitetura do seu processador)
	var numero int64 = 10000
	fmt.Println(numero)

	// Temos também o uint que é um int sem sinal
	var numero2 uint64 = 45555
	fmt.Println(numero2)

	// alias (apelido)
	// no caso do int32
	var numero3 rune = 1235
	fmt.Println(numero3)

	// -- Números reais
	// float32, float64
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	// -- Strings
	// Sempre aspas duplas para strings
	var str string = "Strings"
	fmt.Println(str)

	str2 := "Strings 2"
	fmt.Println(str2)

	// Valor 0
	var texto string
	fmt.Println(texto)
	// vai ser string em branco
	// se for int vai aparecer o número 0

	// Booleano
	var booleano1 bool = true // ou false
	fmt.Println(booleano1)

	// Error
	var erro error
	fmt.Println(erro) // nil

	var erro2 error = errors.New("Esse é o erro!")
	fmt.Println(erro2)
}
