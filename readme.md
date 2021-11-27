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
  - in a more proper setting/if we needed some sort of specific response, we could use postman, but in this case we only effectively need an example target to send a HTTP request to

### App plan

NOTE: This section was written *before* writing the app.  The Overview section will contain something similar but will be updated throughout and after writing the app.

The app will have 4 primary functions as stated above, these will be the 4 primary functions of the app.  The main function will provide the control flow and arguments.

- retrieve(path string) => struct
  - if path not supplied, then default to `./input.json`
  - turns JSON object into a go struct and returns the struct
- validate(data struct) => bool
  - fetches the rules from the schema
  - *ASSUMPTION: we only need to validate according to one particular schema.  If we needed to switch schemas, we could abstract this function by adding an extra schema parameter, or even require the JSON object itself to declare its schema, and fetch it at runtime*
  - validates the struct according to the schema
  - *ASSUMPTION: we allow through extra fields*
  - *ASSUMPTION: all fields listed in the schema are required*
  - returns whether or not the `json` argument is valid according to the schema
- log(data struct) => no return
  - goes through the data structure and logs each part according to type
  - does not return anything
- post(data struct) => API response
  - converts the struct back into a json object
  - sends a POST request to an API including the json object
  - returns the response from the API (which is ultimately discarded, but can be used if the app is extended further)

### Data structure

The JSON data structure will represent a (hypothetical) report on the number of birds present in an area, structured as follows:

```js
{
    "timestamp": 1637980031800,
    "location": [
        {
            "latitude": 0.0,
            "longitude": 0.0
        },
        ...
    ],
    "species": [
        {
            "name": "duck",
            "scientific_name": "Duccus duccus",
            "count": 10,
        }, 
        ...
    ],
}
```

## Overview

The app will be in the following context:

- a separate app provides a JSON object as input to this app, this object will be supplied in the app directory as `input.json`
- we are programming an app that validates this input and logs it before passing it on (unchanged) to be processed by the API located at `https://random-birds.free.beeceptor.com/birdcounts` (via a `POST` method)
- the app will simply be run when necessary (as opposed to being hosted as an always-on service)

## Process

- After planning and writing an example json...
- Scaffolding the app
  - adding the 4 functions and the basic control flow in main()
  - typing the 4 functions and passing arguments in main
    - simple struct for Report to start with
    - will leave the actual return type of post function until we add the http package
- Adding the JSON schema
  - generated from the example json using www.jsonschema.net
  - edited to remove unnecessary details

### Challenges

### Considerations

### Requirements
