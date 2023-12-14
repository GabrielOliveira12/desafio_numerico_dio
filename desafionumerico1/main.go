package main

import "fmt"

func main() {

	var dividendo int = 0
	var divisor int = 3

	for i := 0; i <= 100; i++ {
		if dividendo%divisor == 0 {
			fmt.Println(dividendo)
		} else {
			fmt.Println(dividendo, "Não pode ser dividido por três")
		}
		dividendo++
	}
}
