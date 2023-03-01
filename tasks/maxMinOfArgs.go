package tasks

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func MaxMinOfArgs() {
	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("Give one or more arguments")
		os.Exit(1)
	}

	var n float64
	err := errors.New("an error")
	k := 1

	for err != nil {
		if k >= len(arguments) {
			fmt.Println("None float args")
			os.Exit(1)
		}

		n, err = strconv.ParseFloat(arguments[k], 64)
		k++
	}
	max, min := n, n

	for i := 1; i < len(arguments); i++ {
		arg, err := strconv.ParseFloat(arguments[i], 64)
		if err != nil {
			continue
		}
		if arg > max {
			max = arg
		} else if arg < min {
			min = arg
		}
	}

	fmt.Println(min, max)
}
