package main

import "fmt"

func main() {
	fmt.Println("Welcome to video on slices")

	var fruitList = []string{}
	fmt.Printf("Type of fruitlist is: %T\n", fruitList)
	fruitList = append(fruitList, "Mango", "Banana", "Peach", "Tomato")

	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)
}
