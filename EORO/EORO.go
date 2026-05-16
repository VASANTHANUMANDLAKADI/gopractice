package main
import ("fmt")

func main() {

	var a int
  
    fmt.Println("Enter a number:")
    fmt.Scanf("%d",&a)
    
  if a%2 == 0 {
    
    fmt.Println("Number is even")

  } else { 
  	fmt.Println("Number is odd")
  }
}
