package main
import "fmt"

type person struct {
	name string
	age int
	job string
	salary int
}

func main() {
	var vas person
	var sam person

	fmt.Println("Enter name, age, job and salary of vas:")
	fmt.Scanln(&vas.name, &vas.age, &vas.job, &vas.salary)

	fmt.Println("Enter name, age, job and salary of sam:")
	fmt.Scanln(&sam.name, &sam.age, &sam.job, &sam.salary)

	fmt.Println("Name: ", vas.name)
    fmt.Println("Age: ", vas.age)
    fmt.Println("Job: ", vas.job)
    fmt.Println("Salary: ", vas.salary)

  
    fmt.Println("Name: ", sam.name)
    fmt.Println("Age: ", sam.age)
    fmt.Println("Job: ", sam.job)
    fmt.Println("Salary: ", sam.salary)
}
