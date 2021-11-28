// Code generated by schema-generate. DO NOT EDIT.

package main

import (
	"bytes"
	"encoding/json"
	"errors"
)

// BirdReport A report of the number of birds at a particular time and location
type BirdReport struct {

	// Coordinates to a point and height (NS) and length (EW) of the area
	Location *Location `json:"location"`

	// An array of bird counts organised by species
	Species []*SpeciesItems `json:"species"`

	// Time in ms since epoch
	Timestamp int `json:"timestamp"`
}

// Coordinate A coordinate consisting of latitude and longitude
type Coordinate struct {

	// latitude of coordinate
	Latitude float64 `json:"latitude"`

	// longitude of coordinate
	Longitude float64 `json:"longitude"`
}

// Location Coordinates to a point and height (NS) and length (EW) of the area
type Location struct {

	// A coordinate consisting of latitude and longitude
	Coordinate *Coordinate `json:"coordinate"`

	// North-South length of the area around the coordinate in km
	Height float64 `json:"height"`

	// East-West length of the area around the coordinate in km
	Length float64 `json:"length"`
}

// SpeciesItems Data for a single bird species
type SpeciesItems struct {

	// number of birds of this species sighted
	Count int `json:"count"`

	// common name of the bird species
	Name string `json:"name"`

	// scientific name of the bird species
	ScientificName string `json:"scientific_name"`
}

func (strct *BirdReport) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Location" field is required
	if strct.Location == nil {
		return nil, errors.New("location is a required field")
	}
	// Marshal the "location" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"location\": ")
	if tmp, err := json.Marshal(strct.Location); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Species" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "species" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"species\": ")
	if tmp, err := json.Marshal(strct.Species); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Timestamp" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "timestamp" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"timestamp\": ")
	if tmp, err := json.Marshal(strct.Timestamp); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *BirdReport) UnmarshalJSON(b []byte) error {
	locationReceived := false
	speciesReceived := false
	timestampReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "location":
			if err := json.Unmarshal([]byte(v), &strct.Location); err != nil {
				return err
			}
			locationReceived = true
		case "species":
			if err := json.Unmarshal([]byte(v), &strct.Species); err != nil {
				return err
			}
			speciesReceived = true
		case "timestamp":
			if err := json.Unmarshal([]byte(v), &strct.Timestamp); err != nil {
				return err
			}
			timestampReceived = true
		}
	}
	// check if location (a required property) was received
	if !locationReceived {
		return errors.New("\"location\" is required but was not present")
	}
	// check if species (a required property) was received
	if !speciesReceived {
		return errors.New("\"species\" is required but was not present")
	}
	// check if timestamp (a required property) was received
	if !timestampReceived {
		return errors.New("\"timestamp\" is required but was not present")
	}
	return nil
}

func (strct *Coordinate) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Latitude" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "latitude" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"latitude\": ")
	if tmp, err := json.Marshal(strct.Latitude); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Longitude" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "longitude" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"longitude\": ")
	if tmp, err := json.Marshal(strct.Longitude); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Coordinate) UnmarshalJSON(b []byte) error {
	latitudeReceived := false
	longitudeReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "latitude":
			if err := json.Unmarshal([]byte(v), &strct.Latitude); err != nil {
				return err
			}
			latitudeReceived = true
		case "longitude":
			if err := json.Unmarshal([]byte(v), &strct.Longitude); err != nil {
				return err
			}
			longitudeReceived = true
		}
	}
	// check if latitude (a required property) was received
	if !latitudeReceived {
		return errors.New("\"latitude\" is required but was not present")
	}
	// check if longitude (a required property) was received
	if !longitudeReceived {
		return errors.New("\"longitude\" is required but was not present")
	}
	return nil
}

func (strct *Location) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Coordinate" field is required
	if strct.Coordinate == nil {
		return nil, errors.New("coordinate is a required field")
	}
	// Marshal the "coordinate" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"coordinate\": ")
	if tmp, err := json.Marshal(strct.Coordinate); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Height" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "height" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"height\": ")
	if tmp, err := json.Marshal(strct.Height); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Length" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "length" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"length\": ")
	if tmp, err := json.Marshal(strct.Length); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Location) UnmarshalJSON(b []byte) error {
	coordinateReceived := false
	heightReceived := false
	lengthReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "coordinate":
			if err := json.Unmarshal([]byte(v), &strct.Coordinate); err != nil {
				return err
			}
			coordinateReceived = true
		case "height":
			if err := json.Unmarshal([]byte(v), &strct.Height); err != nil {
				return err
			}
			heightReceived = true
		case "length":
			if err := json.Unmarshal([]byte(v), &strct.Length); err != nil {
				return err
			}
			lengthReceived = true
		}
	}
	// check if coordinate (a required property) was received
	if !coordinateReceived {
		return errors.New("\"coordinate\" is required but was not present")
	}
	// check if height (a required property) was received
	if !heightReceived {
		return errors.New("\"height\" is required but was not present")
	}
	// check if length (a required property) was received
	if !lengthReceived {
		return errors.New("\"length\" is required but was not present")
	}
	return nil
}

func (strct *SpeciesItems) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Count" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "count" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"count\": ")
	if tmp, err := json.Marshal(strct.Count); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Name" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "name" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"name\": ")
	if tmp, err := json.Marshal(strct.Name); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "ScientificName" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "scientific_name" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"scientific_name\": ")
	if tmp, err := json.Marshal(strct.ScientificName); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *SpeciesItems) UnmarshalJSON(b []byte) error {
	countReceived := false
	nameReceived := false
	scientific_nameReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "count":
			if err := json.Unmarshal([]byte(v), &strct.Count); err != nil {
				return err
			}
			countReceived = true
		case "name":
			if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
				return err
			}
			nameReceived = true
		case "scientific_name":
			if err := json.Unmarshal([]byte(v), &strct.ScientificName); err != nil {
				return err
			}
			scientific_nameReceived = true
		}
	}
	// check if count (a required property) was received
	if !countReceived {
		return errors.New("\"count\" is required but was not present")
	}
	// check if name (a required property) was received
	if !nameReceived {
		return errors.New("\"name\" is required but was not present")
	}
	// check if scientific_name (a required property) was received
	if !scientific_nameReceived {
		return errors.New("\"scientific_name\" is required but was not present")
	}
	return nil
}
