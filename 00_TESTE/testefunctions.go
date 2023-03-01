package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func subtrair(n1 int8, n2 int8) int8 {
	return n1 - n2
}

func multiplicar(n1 int8, n2 int8) int8 {
	return n1 * n2
}

func dividir(n1 int8, n2 int8) int8 {
	return n1 / n2
}

func main() {
	soma := somar(30, 10)
	subtrai := subtrair(50, 10)
	multiplica := multiplicar(10, 4)
	dividi := dividir(80, 2)

	fmt.Println(soma)
	fmt.Println(subtrai)
	fmt.Println(multiplica)
	fmt.Println(dividi)

	var f = func() {
		fmt.Println("Função f")
	}

	f()

	var j = func(txt string) {
		fmt.Println(txt)
	}

	j("Função j")

	var i = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := i("Retorno da função i")
	fmt.Println(resultado)
}
