# Quake Data Collector ⚡️

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

You'll need [Git](https://git-scm.com) and [Golang](https://go.dev/doc/install).


## How To Use 

From your command line, clone and run Quake Data Collector:

```bash
# Clone this repository
git clone https://github.com/renatocosta/quake-data-collector-golang.git

# Go into the repository
cd quake-data-collector-golang

## Unit testing
```
go test ./...

## Let's Run the Application 
```
go run src/context/log_handler/main.go
```

## Event Storming

Go through all of the learning journey using Event Storming for understanding the business needs as shown below

### Steps
![Image](./assets/EventStorming.jpg?raw=true)

## Bounded contexts
![Image](./assets/EventStormingOutcome.jpg?raw=true)

[LogHandler](src/Domains/Context/LogHandler)

[MatchReporting](src/Domains/Context/MatchReporting)