package main

import "fmt"

func main() {
    var n int

    fmt.Println("Enter the number to print the multiplication table: ")
    
    fmt.Scanf("%d", &n)

    fmt.Printf("Multiplication Table for %d:\n", n)

    for i := 1; i <= 10; i++ {
        fmt.Printf("%d x %d = %d\n", n, i, n*i)
    }
}
