# Project Assignment

Task: build a Go app to:

- retrieve input from structured JSON
- validate input
- log input to STDOUT
- send a request containing the input to a HTTP based API

## Planning

### Structure / setup

- We will need a repo to host the go app
- We will need ... this planning document to document thought processes and write down the plan
- We will need a JSON structure with fields that require validation
  - Provided in the [Data structure](#data-structure) section
  - We may want to try and create a JSON schema from this
- We will need a JSON input object
  - provided as `input.json` in overview
- We will need a mock API to post the data to
  - found a free service to set up a quick live mock API using Beeceptor

### App plan

### Data structure

The JSON data structure will represent a (hypothetical) report on the number of birds present in an area, structured as follows:

```json
{
    // int: time in ms since epoch
    "timestamp": 1637980031800,
    // coordinates to a quadrilateral enclosed area, must have exactly 4 coordinates
    "location": [
        {
            "latitude": 0.0, // float
            "longitude": 0.0 // float
        },
        // ...
    ],
    // list of species and reported numbers
    "species": [
        {
            // string: common name of the bird species
            "name": "duck",
            // string: scientific name of the bird species
            "scientific_name": "Duccus duccus",
            // int: number of birds
            "count": 10,
        }, 
        // ...
    ],
}
```

## Overview

The app will be in the following context:

- a separate app provides a JSON object as input to this app, this object will be supplied in the app directory as `input.json`
- we are programming an app that validates this input and logs it before passing it on (unchanged) to be processed by the API located at `https://random-birds.free.beeceptor.com/birdcounts`
- the app will simply be run when necessary (as opposed to being hosted as an always-on service)

## Process

### Challenges

### Considerations

### Requirements
