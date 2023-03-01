package tasks

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"regexp"
)

func findIP(input string) string {
	numBlock := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	regexPattern := numBlock + "\\." + numBlock + "\\." + numBlock + "\\." + numBlock

	regEx := regexp.MustCompile(regexPattern)
	return regEx.FindString(input)
}

func FindIpV4() {
	args := os.Args

	if len(args) < 2 {
		fmt.Printf("usage: %s logFile\n", filepath.Base(args[0]))
		os.Exit(1)
	}

	for _, filename := range args[1:] {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		defer file.Close()
		fmt.Println("Ip from file: ", file.Name())
		reader := bufio.NewReader(file)

		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("error reading file: %s", err)
				break
			}

			ip := findIP(line)
			trial := net.ParseIP(ip)
			if trial.To4() == nil {
				continue
			}
			fmt.Println("ip:", ip)
		}
	}
}
