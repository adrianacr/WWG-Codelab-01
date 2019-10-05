package main

import "fmt"

func main() {
	println("exercicio 01 do lab 1 :D")
	var primeiroNumero int
	var segundoNumero int
	var terceiroNumero int
	primeiroNumero = 15
	segundoNumero = 30
	terceiroNumero = 45

	fmt.Println("Primeiro numero", primeiroNumero)
	fmt.Println("Segundo numero", segundoNumero)
	fmt.Println("Terceiro numero", terceiroNumero)

	println("\nexercicio 02 do lab 1 :D")

	a := 230
	b := 27

	soma := a + b
	subtração := a - b

	fmt.Println("A variavel A vale", a)
	fmt.Println("A variavel B vale", b)
	fmt.Println("A variavel SOMA vale", soma)
	fmt.Println("A variavel SUBTRACAO vale", subtração)

	println("\nexercicio 03 do lab 1 :D")

	multiplicação := a * b
	divisão := a / b

	fmt.Println("A variavel MULTIPLICAÇÃO vale", multiplicação)
	fmt.Println("A variavel DIVISÃO vale", divisão)

	println("\nexercicio 04 do lab 1 :D")

	fmt.Printf("\nOs tipos de A e B são, respectivamente: %T e %T", a, b)
	fmt.Printf("\nOs tipos de somaAB e subtraçãoAB são, respectivamente: %T e %T", soma, subtração)
	fmt.Printf("\nOs tipos de multiplicaçãoAB e divisãoAB são, respectivamente: %T e %T\n", multiplicação, divisão)

	println("\nexercicio 05 do lab 1 :D")

	var preçoDoPão float64 = 4.60
	var preçoDaAveia float64 = 5
	var preçoDoAzeite float64 = 19.95
	var preçoDoSuco float64 = 7
	var preçoDaÁgua float64 = 2.15
	var preçoDoKGMaçã float64 = 8.90
	var preçoDoKGBanana float64 = 4.5

	valorDaCompra := (preçoDoPão * float64(3)) + (preçoDaAveia * float64(0)) + (preçoDoAzeite * float64(1)) + (preçoDoSuco * float64(2)) + (preçoDaÁgua * float64(5)) + (preçoDoKGMaçã * float64(1.5)) + (preçoDoKGBanana * float64(0))

	fmt.Println("O valor da compra foi", valorDaCompra)

	println("\nexercicio 06 do lab 1 :D")

	fmt.Println("var é uma palavra reservada, não pode ser usada para nomear uma variável")

	println("\nexercicio 07 do lab 1 :D")

	X := 15
	var Y int
	Y = 31
	var Z = 47

	fmt.Println("\n\t X \t Y \t Z")
	fmt.Printf("base 2 : %b \t%b \t%b", X, Y, Z)
	fmt.Printf("\nbase 10: %d \t%d \t%d", X, Y, Z)
	fmt.Printf("\nbase 16: %#X \t%#X \t%#X\n", X, Y, Z)

	println("\nexercicio 08 do lab 1 :D")

	var w int

	fmt.Printf("a variavel W tem valor: %b", w)
	fmt.Printf("A VARIÁVEL RECEBE VALOR 0 AUTOMATICAMENTE")

}
