package main

import "fmt"

type Employee interface {
	Language() string
	Age() int
}

type Engineer struct {
	Name        string
	EngineerAge int
}

func (e Engineer) Language() string {
	return e.Name + " programs in Go"
}

func (e Engineer) Age() int {
	return e.EngineerAge
}

func main() {

	var programmers []Employee
	elliot := Engineer{Name: "Elliot", EngineerAge: 35}
	mirquo := Engineer{Name: "Mirquo", EngineerAge: 36}

	programmers = append(programmers, elliot)

	fmt.Printf("%d is the lengh and %d is the cap\n", len(programmers), cap(programmers))

	programmers = append(programmers, mirquo)

	fmt.Printf("%d is the lengh and %d is the cap\n", len(programmers), cap(programmers))
	programmers = programmers[1:2]

	fmt.Printf("%d is the lengh and %d is the cap\n", len(programmers), cap(programmers))

	fmt.Printf("The programmers name is %s, age %d and %s\n", elliot.Name, elliot.EngineerAge, elliot.Language())

	maks := CreateEngineer("Maks", 29)

	programmers = AppendProgrammer(programmers, maks)

	fmt.Printf("%d is the lengh and %d is the cap\n", len(programmers), cap(programmers))

	max := CreateEngineer("Max", 28)

	AppendProgrammerPointer(&programmers, max)

	fmt.Printf("%d is the lengh and %d is the cap\n", len(programmers), cap(programmers))

	for _, programmer := range programmers {
		engineer, ok := programmer.(Engineer)
		if ok {
			fmt.Printf("The programmer's name is %s, age %d and %s\n", engineer.Name, engineer.EngineerAge, engineer.Language())
		} else {
			fmt.Println("Type assertion failed")
		}
	}

}

func CreateEngineer(name string, age int) Engineer {
	return Engineer{Name: name, EngineerAge: age}
}

func AppendProgrammer(em []Employee, en Engineer) []Employee {
	em = append(em, en)
	return em

}

func AppendProgrammerPointer(em *[]Employee, en Engineer) {
	*em = append(*em, en)
}
