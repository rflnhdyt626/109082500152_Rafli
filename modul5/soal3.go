package main

import "fmt"

func printFaktor(n int, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	printFaktor(n, i+1)
}

func main() {
	var n int
	fmt.Scanln(&n)
	printFaktor(n, 1)
	fmt.Println()
}
