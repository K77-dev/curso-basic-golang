package main

import (
	"fmt"
	c "strconv"
)

func main(){
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))
	fmt.Println(int(x) / y)

	// cuidado... imprimirá a letra 'a'
	//fmt.Println("Teste", string(97))

	// int para string
	fmt.Println("Teste", c.Itoa(97))

	// string para int
	num, _ := c.Atoi("97") // num é a variável de retorno e _ ó o erro
	fmt.Println(num)
	fmt.Println(num - 10)

	b, _ := c.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}
}