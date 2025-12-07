package main

import "fmt"

func main(){
	// var aprovados map[int]string // erro porque map tem que ser inicializados
	aprovados := make(map[int]string)

	aprovados[12345678909] = "Maria"
	aprovados[18765432123] = "João"
	aprovados[77600274191] = "Kelsen"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%d %s\n", cpf, nome)
	}

	fmt.Println(aprovados[77600274191])

	delete(aprovados, 18765432123)
	for _, nome := range aprovados {
		fmt.Printf("%s\n", nome)
	}
}