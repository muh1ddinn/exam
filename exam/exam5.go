package main

import (
	"fmt"
	"os"
)

type Student struct {
	name   string
	grades []int
	course string
}

func (s Student) calculateAverage() float64 {
	sum := 0
	for _, grade := range s.grades {
		sum += grade
	}
	return float64(sum) / float64(len(s.grades))
}

func writeResultToFile(filename string, data string) {
	err := os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		fmt.Printf("err whil writing to file: %v", err)
		return
	}
}

func main() {
	student := Student{
		name:   "I ",
		grades: []int{70, 80, 90, 100},
		course: "Computer Science",
	}

	average := student.calculateAverage()
	var result string
	if average < 60 {
		result = "You are failed"
	} else {
		result = "You are passed"
	}

	writeResultToFile("result.txt", result)
}
