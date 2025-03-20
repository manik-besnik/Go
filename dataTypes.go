package main

import "fmt"

func main() {
	stringDataType()
	intDataType()
	uintDataType()
	fixSizeArr()
	dynamicSizeArr()
	mapDataType()
	structType()

	condition()

	var s = Rectangle{width: 10, height: 5}
	fmt.Printf("Rectangle Area: %.2f\n", s.Area())
}

func intDataType() {
	var intData int = -100
	fmt.Println(intData)
}

func uintDataType() {

	var intData uint = 1100
	fmt.Println(intData)
}

func stringDataType() {
	var name string = "Manik"
	fmt.Printf("My Name is: %s \n", name)
}

func fixSizeArr() {
	var fixArr [3]int = [3]int{10, 20, 30}

	for index, value := range fixArr {
		fmt.Printf("Array Index %d and array value %d \n", index, value)
	}
}

func dynamicSizeArr() {
	var fixArr []int = []int{10, 20, 30, 40}

	for index, value := range fixArr {
		fmt.Printf("Dymanic Array Index %d and array value %d \n", index, value)
	}
}

func mapDataType() {
	var m map[string]interface{} = map[string]interface{}{"name": "manik", "age": 25}

	for key, value := range m {
		fmt.Printf("Map Key: %s, Map Value: %v\n", key, value)
	}
}

type Person struct {
	name string
	age  int
}

func structType() {
	var person = Person{name: "Manik", age: 25}
	fmt.Printf("Name: %s, Age: %d\n", person.name, person.age)
}

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func condition() {
	number := 80

	if number > 79 {
		fmt.Print("Grade is A+ \n")
	} else if number > 69 {
		fmt.Print("Grade is A \n")
	} else if number > 59 {
		fmt.Print("Grade is A- \n")
	} else if number > 49 {
		fmt.Print("Grade is B \n")
	} else if number > 39 {
		fmt.Print("Grade is C \n")
	} else if number >= 33 {
		fmt.Print("Grade is D \n")
	} else {
		fmt.Print("Failed")
	}
}
