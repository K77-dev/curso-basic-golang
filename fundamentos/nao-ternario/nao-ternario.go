package main

import "fmt"

// não há operador ternário

func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
	// return nota >=6 ? "Aprovado" : "Reprovado" // não existe 
}

func main(){
	fmt.Println(obterResultado(7))
}
