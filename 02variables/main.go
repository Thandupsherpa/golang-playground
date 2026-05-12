package main

import "fmt"

// variables, types, and constants

//const name := "thandup" - walrus operator can only be used inside a  function not  in global scope

func main() {

	fmt.Println("Variables in Golang")

	// basic variable declaration

	//string
	var name string = "Thandup"
	fmt.Println("Hello", name)

	//int
	var amount int = 6969
	fmt.Println("Amount", amount)

	//boolean
	var isLoggedIn bool = true
	fmt.Println("is user looged in:", isLoggedIn)

	//float
	var weight float32 = 56.5
	fmt.Println("weight is:", weight)

	// checking types
	var food string = "momo"
	fmt.Printf("food type is %T \n", food)

	//with walrus
	// the := operator automatically detects data type of age(int)
	//:= cannot be used outside a method
	age := 19
	fmt.Println("Your age is", age)

	//without type
	// automatically detects the type looking at the value
	var surname = "sherpa"
	fmt.Println(surname)
	fmt.Printf("surname type  is %T \n", surname)

	//constants

	// value can be reassigned  with var
	// allowed
	var year int = 2025
	year = 2626
	fmt.Println(year)

	// uint can only store 0 and positive numbers
	var positive uint = 10
	var negative int =  -10
	//  var nev uint = -10 err
	fmt.Println(positive,negative)

	const salary int = 696969
	//salary = 969696 - Not allowed
	fmt.Println(salary)
}
