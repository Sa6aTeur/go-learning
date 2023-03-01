package tasks

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func SumWhileNotEnd() {
	fmt.Println("Please enter number for sum")
	var sum float64
	file := os.Stdin
	defer file.Close()

	scanner := bufio.NewScanner(file)

	fmt.Print("Number: ")
	for scanner.Scan() {

		if scanner.Text() == "end" {
			break
		}
		value, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("Please enter number")
			fmt.Print("Number: ")
		} else {
			sum += value
			fmt.Println("Sum: ", sum)
			fmt.Print("Number: ")
		}
	}

}
