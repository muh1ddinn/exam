package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	numbersStr := strings.Split(string(data), ",")

	numbersFile, err := os.Create("numbers.txt")
	if err != nil {
		log.Fatalf("Error creating numbers.txt: %v", err)
	}
	defer numbersFile.Close()

	var totalSum int

	for _, numStr := range numbersStr {
		numStr = strings.TrimSpace(numStr)

		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Printf("Error converting string to integer: %v\n", err)
			continue
		}

		if _, err := numbersFile.WriteString(numStr + "\n"); err != nil {
			log.Fatalf("Error writing to numbers.txt: %v", err)
		}

		totalSum += num
	}

	if _, err := numbersFile.WriteString(fmt.Sprintf("Total sum: %d\n", totalSum)); err != nil {
		log.Fatalf("Error writing total sum to numbers.txt: %v", err)
	}

	fmt.Println("Files written successfully.")
}
