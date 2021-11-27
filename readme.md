# Project Assignment

Task: build a Go app to:

- retrieve input from structured JSON
- validate input
- log input to STDOUT
- send a request containing the input to a HTTP based API

## Planning

### Data structure

The JSON data structure will represent a (hypothetical) report on the number of birds present in an area, structured as follows:

```json
{
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

## Process

### Challenges

### Considerations

### Requirements
