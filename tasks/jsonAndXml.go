package tasks

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadAndWriteJson() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Choose one more file")
		return
	}

	file := args[1]
	fileData, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	parsedData := make(map[string]interface{})

	json.Unmarshal([]byte(fileData), &parsedData)

	for key, value := range parsedData {
		fmt.Println(key, ": ", value)
	}
	rec, err := json.Marshal(parsedData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(rec))
}

func ReadAndWriteXML() {

}
