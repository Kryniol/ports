# Port data parser

Author: Pawel Krynicki <pawel.krynicki@hotmail.com>

## Overview

This CLI application allows saving port data provided as a JSON file. Only in-memory storage is available.

Assumptions:

* `coordinates` are provided in reverse order so that the longitude comes first, the field itself is optional
* `timezone` follows IANA timezone names
* `regions` field contains a list of strings
* `unlocs` field contains a list of port IDs
* invalid data should not cause the whole process to be terminated

Future improvements:

* introduce acceptance/black-box tests based on Docker image
* introduce some kind of persistent storage
* consider implementing command pattern for saving ports and to move the json data reader to the application layer

## Architecture

The app is divided into following layers:

* domain - storing the core domain logic and business rules encapsulated in entities, services and value objects
* infrastructure - providing low-level adapters for ports defined in the domain

## Requirements

Go: 1.17

## Running the app

The binary is included - it's compiled for amd64 architecture and Linux OS. In order to run the CLI application enter
the root directory of the repository and execute in your terminal:

```shell script
./cli -path={path_to_input_file}
```

## Building from sources

To build the binary run the following:

```shell script
go mod download
go build -o cli cmd/cli/main.go
```

Note that you need Go 1.17 in order to build the app from sources.

## Run with Docker
Due to poor internet connection (I'm on vacation right now) I couldn't fully test the Docker setup.
In order to run the app in a container try executing the following commands:
```shell script
docker build -t ports_cli .
docker run --rm -v {absolute_path_to_input_file}:/app/input.json ports_cli /app/cli -path /app/input.json
```

## Running the tests

The codebase is covered by unit and integration tests. In order to run the whole test suite execute:

```shell script
go mod download
go test -v ./...
```
