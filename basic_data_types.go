package main

import "fmt"

func main() {
	var name string = "Mustafa"
	var age int = 40
	var job bool = true
	
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(job)
	

	fmt.Printf("%T\n", name) 
	fmt.Printf("%T\n", job)
	
}
