package main

import (
	"fmt"
	"mylearning/myutil"
)

func main() {
	fmt.Println("My name is Dixita")

	myutil.PrintMessage("Hello everyone")

	// var name string = "Dixita"
	// var version = "latest version"
	// // if not described any var type then compiler understands it automatically at compile-time
	// fmt.Println(name)
	// fmt.Println(version)

	// var money int = 67000
	// var currency = "USD"
	// fmt.Println(money)
	// fmt.Println("this is my saved amount:", money, currency)

	// var dimension float64 = 87.12
	// fmt.Println(dimension)

	// var decided bool = false
	// fmt.Println(decided)

	// var person = "Dixita"
	// fmt.Println(person)

	// const pi = 67.12
	// fmt.Println(pi)

	person := 123
	fmt.Println(person)

	var Public = "data is important"
	var private = "data is important"

	fmt.Println(Public)
	fmt.Println(private)

	//Public function(exported)
	// func PublicFunction(){}

	//Private function (unexported)
	// func privateFunction(){}
}
