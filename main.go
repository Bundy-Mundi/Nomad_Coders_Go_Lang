package main

import (
	"fmt"

	"github.com/Bundy-Mundi/learnGo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"apple": "A red colored fruit", "pear": "A delicious fruit"}
	err := dictionary.Add("apples", "Two Apples")
	if err == nil {
		fmt.Println("Successfully Added to the dictionary")
	} else {
		fmt.Println(err)
	}
	dictionary.Update("apples", "New UPDATED apples")
	dictionary.Add("Chipe", "So Cheap")
	v, err2 := dictionary.Delete("pear")
	if err2 == nil {
		command := fmt.Sprintf("%s Has Been Deleted Completely", v)
		fmt.Println(command)
	} else {
		fmt.Println(err2)
	}
	fmt.Println(dictionary)
}
