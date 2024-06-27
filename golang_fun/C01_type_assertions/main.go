package main

import (
	"fmt"
)

type Developer struct {
	Name string
	Age  int
}

func GetDeveloper(name interface{}, age interface{}) Developer {
	nameStr, ok := name.(string)
	if !ok {
		fmt.Println("Varianle one is not a string.")
		nameStr = "Unknow"
	}
	ageInt, ok := age.(int)
	if !ok {
		fmt.Println("Variable two is not an integer.")
		ageInt = 0
	}

	return Developer{
		Name: nameStr,
		Age:  ageInt,
	}

}

func main() {
	fmt.Println("Hello World")

	var name interface{} = "Elliote"
	var age interface{} = 26

	dynamicDev := GetDeveloper(name, age)
	fmt.Println(dynamicDev.Name)
}
