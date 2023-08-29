package main

import "fmt"

func main (){
	fmt.Println("Welcome to arrays in golang")

	var fruitList [5] string

	fruitList[0] = "Pear"
	fruitList[1] = "Banana"
	fruitList[2] = "Peach"
	fruitList[4] = "Kiwi"
	fmt.Println("Fruit list is : ", fruitList)
	fmt.Println("Fruit list is : ", len(fruitList))


	var vegList = [5] string{"onion", "potato"}
	fmt.Println("Veg list is : ", vegList)

}