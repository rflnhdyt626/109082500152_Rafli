package main

import "fmt"

func printBarisan(n int) {
	if n == 0 {
		return
	}
	fmt.Print(n, " ")
	printBarisan(n - 1)
	fmt.Print(n, " ")
}

func main() {
	var n int
	fmt.Scanln(&n)
	printBarisan(n)
}