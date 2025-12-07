package main

import "fmt"

func main(){
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Giovanna Brito": 20600.00,
			"Gabriela Silva": 14500.00,
		},
		"H": {
			"Hélio Rubens": 23233.33,
		},
		"K": {
			"Kelsen Brito": 23423.00,
			"Kerginaldo Brito": 2334.34,
		},
	} 

	delete(funcsPorLetra, "H")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
		for i, v := range funcs {
			fmt.Println(i, v)
		}
		fmt.Printf("\n")
	}
}