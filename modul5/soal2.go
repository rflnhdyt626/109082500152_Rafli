package main

import "fmt"

func printBintang(n int) {
	if n <= 0 {
		return
	}
	printBintang(n - 1)

	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Scanln(&n)
	printBintang(n)
}
