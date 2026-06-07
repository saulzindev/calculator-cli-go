package main

import (
	"fmt"
	"math"
)

func main() {
	var n1, n2 int
	var escolha int

	fmt.Println("---------BEM VINDO A CALCULADORA SAUL--------")

	fmt.Print("Digite um numero: ")
	fmt.Scanln(&n1)

	fmt.Print("Digite outro numero: ")
	fmt.Scanln(&n2)

	fmt.Println("----------ESCOLHA SUA OPERAÇÃO----------")
	fmt.Println("1- Soma")
	fmt.Println("2- Subtração")
	fmt.Println("3- Multiplicação")
	fmt.Println("4- Divisão")
	fmt.Println("5- Raiz Quadrada")
	fmt.Scanln(&escolha)

	switch escolha {
	case 1:
		fmt.Println("Resultado:", n1+n2)
	case 2:
		fmt.Println("Resultado:", n1-n2)
	case 3:
		fmt.Println("Resultado:", n1*n2)
	case 4:
		if n2 == 0 {
			fmt.Println("Não é possível dividir por zero.")
		} else {
			fmt.Println("Resultado:", float64(n1)/float64(n2))
		}
	case 5:
		fmt.Println("Raiz de", n1, "=", math.Sqrt(float64(n1)))
		fmt.Println("Raiz de", n2, "=", math.Sqrt(float64(n2)))
	default:
		fmt.Println("Opção inválida")
	}
}
