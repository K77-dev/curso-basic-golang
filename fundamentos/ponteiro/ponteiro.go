package main

import "fmt"

func main(){
	i := 1

	// Go não tem aritimética de ponteiros
	var p *int = nil
	p = &i // pegando o endereço da variável i
	fmt.Println(p, *p, &i, i)
	
	*p++ // desreferenciando (pegando o valor)
	i++
	fmt.Println(p, *p, &i, i)

	var q *int = nil
	q = p // ou &i (endereço de i)
	fmt.Println(q, *q)

	// p++ não pode porque p é um ponteiro para int e não um int
}