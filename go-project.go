package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func retrieveData(path string) BirdReport {
	if path == "" {
		path = "./input.json"
	}

	// open json file
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var report BirdReport
	err = report.UnmarshalJSON(byteValue)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("successfully unmarshalled JSON")

	return report
}

func validateData(data BirdReport) bool {
	return true
}
func logData(data BirdReport) {
}
func postData(data BirdReport) string {
	return ""
}

func main() {
	data := retrieveData("")
	valid := validateData(data)
	if !valid {
		fmt.Println("Input JSON was not valid")
		os.Exit(1)
	}
	logData(data)
	postData(data)
}
