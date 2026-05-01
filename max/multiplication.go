package main

import "fmt"

func main2() {

    var n int

    fmt.Print("Enter the number to print the multiplication table: ")

    fmt.Printf("\nMultiplication Table for %d:\n", n)

    for i := 1; i <= 10; i++ {
        fmt.Printf("%d x %d = %d\n", n, i, n*i)
    }
}
