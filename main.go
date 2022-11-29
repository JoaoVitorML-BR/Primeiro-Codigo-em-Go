package main

import (
	// "bufio"
	"fmt"
	// "os"
	// "strings"
)

func menu() {
    
	fmt.Println("Qual operação deseja fazer?")

	fmt.Println("1) Soma ")
	fmt.Println("2) Subtração ")
	fmt.Println("3) Multiplicação ")
	fmt.Println("4) Divisão")


}

func main() {
	var num1, num2, resposta int
	//soma e subtracao tem que ir para depois de obter as entradas

	fmt.Println("Bem-vindo(a) a calculadora!")

	fmt.Println("Digite o valor do primeiro número: ")
	fmt.Scanln(&num1) // aqui deveria ser num1

	fmt.Println("Digite o valor do segundo número: ")
	fmt.Scanln(&num2)

	menu()
	fmt.Scanln(&resposta)
	switch resposta {
	case 1:
		var soma int = num1 + num2	
		fmt.Println(soma)
	case 2:
			var subtracao int = num1 - num2
			fmt.Println(subtracao)
	case 3:
		var multiplicacao int = num1 * num2 
		fmt.Println(multiplicacao)
	case 4:
		var divisao int = num1 / num2
		fmt.Println(divisao)
	default:
			fmt.Println("Número inválido!")
	}

}