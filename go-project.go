package main

import (
	"fmt"
	"os"
)

type Report struct {
	Timestamp int64
}

func retrieveData(path string) Report {
	if path == "" {
		path = "./input.json"
	}
	return Report{}
}

func validateData(data Report) bool {
	return false
}
func logData(data Report) {
}
func postData(data Report) string {
	return ""
}

func main() {
	var data Report = retrieveData("")
	var valid bool = validateData(data)
	if !valid {
		fmt.Println("Input JSON was not valid")
		os.Exit(1)
	}
	logData(data)
	postData(data)
}
