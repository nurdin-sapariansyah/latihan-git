package main

import (
	"fmt"
)


func main(){
	i := 0
	fmt.Print(i)
	fmt.Println("---------->>Identitas saya<<---------------")
	//fiturA
	name := "M. nurdin sapariansyah"
	age := 17

	fmt.Println("nama saya adalah: ", name)
	fmt.Println("umur saya adalah: ", age)
	fmt.Println("---------->>Background saya<<---------------")
	//fiturB
	for i:=0; i<=5; i++{
		for j:=0; j<i; j++{
			print("a")
		} 
		println("b")
	}
	
}