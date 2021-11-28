package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func retrieveData(path string, report *BirdReport) error {
	if path == "" {
		path = "./input.json"
	}

	// open json file
	jsonFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	// unmarshal json
	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = report.UnmarshalJSON(byteValue)
	if err != nil {
		return err
	}
	return nil
}

func validateData(data BirdReport) error {
	// this way we can warn the user of all errors at once rather than one at a time
	errs := ""

	// validate loc.coord.long and lat
	loc := data.Location
	coord := loc.Coordinate
	if coord.Latitude < -90 || coord.Latitude > 90 {
		errs += fmt.Sprintf("invalid latitude: %v\n", coord.Latitude)
	}
	if coord.Longitude < -180 || coord.Longitude > 180 {
		errs += fmt.Sprintf("invalid longitude: %v\n", coord.Longitude)
	}
	// validate loc.len and height
	if loc.Length <= 0 {
		errs += fmt.Sprintf("invalid length: %v\n", loc.Length)
	}
	if loc.Height <= 0 {
		errs += fmt.Sprintf("invalid height: %v\n", loc.Height)
	}
	// validate species counts
	for _, species := range data.Species {
		if species.Count < 0 {
			errs += fmt.Sprintf("invalid species count: %s: %v\n", species.Name, species.Count)
		}
	}

	if errs != "" {
		return errors.New(errs)
	}
	return nil
}

func logData(data BirdReport) {
	fmt.Println()
	fmt.Println("logging bird report data:")
	// log timestamp
	timestamp := time.Unix((int64)(data.Timestamp), 0)
	fmt.Printf("timestamp: %s\n", timestamp.String())
	fmt.Println()

	// log location
	fmt.Println("location:")
	fmt.Printf("coordinates: %v, %v\n", data.Location.Coordinate.Latitude, data.Location.Coordinate.Longitude)
	fmt.Printf("length*height (km): %v * %v\n", data.Location.Length, data.Location.Height)
	fmt.Println()

	// log bird counts
	fmt.Println("bird counts:")
	for _, species := range data.Species {
		fmt.Printf("%s (%s): %v\n", species.Name, species.ScientificName, species.Count)
	}
	fmt.Println()
}

func postData(data BirdReport) string {
	return ""
}

func main() {
	var report BirdReport

	err := retrieveData("", &report)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("successfully unmarshalled JSON")

	err = validateData(report)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("successfully validated JSON")

	logData(report)
	postData(report)
}
