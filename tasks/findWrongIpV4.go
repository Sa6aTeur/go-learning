package tasks

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func findWriteIP(input string) bool {
	numBlock := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	regexPattern := numBlock + "\\." + numBlock + "\\." + numBlock + "\\." + numBlock

	regEx := regexp.MustCompile(regexPattern)
	return regEx.MatchString(input)
}

func FindWrongIpV4() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Choose one more file")
		os.Exit(2)
	}

	for _, filename := range args[1:] {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer file.Close()
		reader := bufio.NewReader(file)

		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Println(err)
				break
			}
			isValidIp := findWriteIP(line)
			if isValidIp {
				fmt.Println(strings.Replace(line, "\n", " : valid ip", -1))
			} else {
				fmt.Println(strings.Replace(line, "\n", " : not valid", -1))
			}
		}
	}
}
