package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	filename := ""
	fmt.Scan(&filename)

	n := 3
	for i := 0; i < n; i++ {
		newFilename := filename + strconv.Itoa(rand.Intn(50))
		if !checkFileExists(newFilename + ".txt") {
			_, err := os.Create(newFilename + ".txt")
			if err != nil {
				fmt.Println("Error creating file:", err)
				return
			}
			fmt.Println("Created file:", newFilename)
			break
		}
	}
}

func checkFileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !errors.Is(err, os.ErrNotExist)
}
