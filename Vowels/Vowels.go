package main
import (
	"fmt" 
  	"strings")

func main(){

	var ip string
	count := 0

	fmt.Println("Enter a word:")
	fmt.Scan(&ip)

	ip = strings.ToLower(ip)

	for _,char := range ip {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u'{
			count++
		}
	}
	fmt.Println("Number of Vowels in the word are :", count)
	
}
