package main

import "fmt"

func main(){
	funcsESalarios := map[string]float64{
		"Kelsen Brito": 20600.00,
		"Gabriela Silva": 14500.00,
		"Pedro Júnior": 5600.10,
	}

	funcsESalarios["Giovanna Brito"] = 34409.90
	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}