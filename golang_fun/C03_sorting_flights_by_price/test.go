package main

import "fmt"

type Programmer struct {
	Name string
	Age  int
}

func (p Programmer) Info() string {
	return fmt.Sprintf("The programmers name is %s and the age is %d", p.Name, p.Age)
}

func NewProgrammer(name interface{}, age interface{}) Programmer {
	nameStr, ok := name.(string)
	if !ok {
		fmt.Println("The name input is not a string. Will input unknown.\n")
		nameStr = "Unknown"
	}
	ageInt, ok := age.(int)
	if !ok {
		fmt.Println("The age input is not an int. Will input \"0\".\n")
		ageInt = 0
	}

	return Programmer{
		Name: nameStr,
		Age:  ageInt,
	}
}

func valueswithc(name *string) string {
	*name = "Switch bijatch"
	return *name
}

func main() {

	mirko := Programmer{
		Name: "Mirko",
		Age:  35,
	}

	maks := NewProgrammer("Makše", "29")

	fmt.Println(mirko.Info())
	fmt.Println(maks.Info())

	nameus := "Tustikus"
	nameusPointer := &nameus

	fmt.Println(&nameus)
	fmt.Println(nameusPointer)
	fmt.Println(&nameusPointer)

	*nameusPointer = "Marko"
	fmt.Println(*nameusPointer)
	fmt.Println(nameus)

	fmt.Println(valueswithc(&nameus))

}
