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
	soma := somar(10, 30)
	subtrai := subtrair(40, 10)
	multiplica := multiplicar(10, 3)
	dividi := dividir(90, 3)

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

	j("Retorno da Função J")

	var i = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := i("Retorno da Função i")
	fmt.Println(resultado)

}
