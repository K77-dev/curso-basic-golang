package main

import "fmt"

func main(){
	// homogênia (mesmo tipo) e estático (tamanho fixo)
	var notas [3]float64
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1
	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))
	fmt.Printf("Média %.2f", media)

	//var p *[len(notas)]float64 = nil // pode ser inicializada assim
	p := &notas
	fmt.Println("")
	for i, v := range p { // i = índice, v = valor
		fmt.Println(i, &p[i], p[i], v)
	}
}