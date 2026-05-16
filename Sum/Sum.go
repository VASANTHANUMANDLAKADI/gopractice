package main
import "fmt"

func main() {
	var N int
	sum := 0

	fmt.Println("Enter a number")
	fmt.Scan(&N)

	for i := 1; i <= N; i++ {
			sum = sum + i
	}

	fmt.Println(sum)
}
