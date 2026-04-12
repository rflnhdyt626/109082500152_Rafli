package main

import "fmt"
func printGanjil(n int) {
	if n <= 0 {
		return
	}
	printGanjil(n - 1)
	if n%2 != 0 {
		fmt.Print(n, " ")
	}
}

func main() {
	var n int
	fmt.Scanln(&n)
	printGanjil(n)
	fmt.Println()
}