package main

import "fmt"

func main() {
	var n, IP, OP int
	
	fmt.Println("Enter a number:")
	fmt.Scanln(&n)

	IP = n
	
	for n > 0 {
	num := n%10
	OP = OP*10 + num
	n = n/10
	}
	
	if IP == OP {
		fmt.Println("Number is a Palindrome")
	} else {
		fmt.Println("Number is not a Palindrome")
		}
}