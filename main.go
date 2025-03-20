package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}

	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	person := map[string]string{
		"name": "John",
		"age":  "30",
	}

	for key, value := range person {
		fmt.Printf("%s: %s\n", key, value)
	}

	learnDataTypes()
}

func learnDataTypes() {
	var x int = 30

	fmt.Println(x)

	var arr [3]int = [3]int{1, 2, 3}

	for index, value := range arr {
		fmt.Printf("Index:%d and value %d \n", index, value)
	}
}
