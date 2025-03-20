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
}
