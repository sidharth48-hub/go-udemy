package main

import "fmt"

func main() {
	age := 32

	agePointer := &age

	fmt.Println("age", *agePointer)
	getAdultYears(agePointer)
	fmt.Println("Adult years: ", age)
}

func getAdultYears(age *int) {
	*age = *age - 18
}
