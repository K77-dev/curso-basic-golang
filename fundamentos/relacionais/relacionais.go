package main

import (
	"fmt"
	"time"
)

func main(){
	banana1 := "Banana"
	banana2 := "Banana"
	fmt.Println("String: ", banana1 == banana2)

	d1 := time.Unix(0,0)
	d2 := time.Unix(0,0)
	fmt.Println("Datas:", d1.Equal(d2))
	
	type Pessoa struct{
		Nome string
	}
	p1 := Pessoa{"Kelsen"}
	p2 := Pessoa{"Kelsen"}
	fmt.Println("Pessoas:", p1 == p2)

}