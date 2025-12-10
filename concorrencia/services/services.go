package main

import (
	"fmt"
	"time"
)

func fetchData(service string, ch chan string) {
	// simula tempo de resposta da API
	time.Sleep(time.Duration(len(service)) * 300 * time.Millisecond)
	ch <- fmt.Sprintf("Response from %s", service)
}

func main() {
	services := []string{"users", "orders", "payments"}

	ch := make(chan string)

	// dispara goroutines
	for _, s := range services {
		go fetchData(s, ch)
	}

	// recebe resultados (um por goroutine)
	for range services {
		result := <-ch
		fmt.Println(result)
	}

	fmt.Println("All services responded!")
}
