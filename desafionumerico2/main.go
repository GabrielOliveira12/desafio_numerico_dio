package main

import "fmt"

func main() {

	for i := 0; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i, "Pin")
		} else {
			fmt.Println(i)
		}
		if i%5 == 0 {
			fmt.Println(i, "pan")
		}
	}
}
