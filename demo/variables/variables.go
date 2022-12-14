package main

import (
	"fmt"
)

func main() {
	var myName = "Sabry"
	fmt.Println("My Name: ", myName)

	var name string = "Sabry"
	fmt.Println("My Name: ", name)

	username := "Sabry"
	fmt.Println("Username: ", username)

	var sum int
	fmt.Println("sum: ", sum)

	part1, other := 1, 5
	fmt.Println("part1: ", part1, "other: ", other)

	part2, other := 2, 5
	fmt.Println("part1: ", part2, "other: ", other)

	sum = part1 + part2
	fmt.Println("sum: ", sum)

	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)

	fmt.Println("lessonName: ", lessonName, "LessonType: ", lessonType)

	word1, word2, _ := "Hello", "World", "!"

	fmt.Println("Word1", word1, "Word2", word2)
}
