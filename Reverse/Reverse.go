package main

import "fmt"

func main(){
	var n int
	rev := 0

	fmt.Println("Enter a number:")
	fmt.Scan(&n)

	for n > 0{
	digit := n%10
	rev = rev*10 + digit
	n = n/10
	}

	fmt.Println("Reversed number", rev)
	
}
