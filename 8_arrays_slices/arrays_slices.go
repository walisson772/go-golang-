package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array [5]string
	array[0] = "Posição 1"
	fmt.Println(array)

	array2 := [5]string{"Posiçao 1", "Posiçao 2", "Posiçao 3", "Posiçao 4", "Posiçao 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slice)

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição alterada"
	fmt.Println(slice2)

}