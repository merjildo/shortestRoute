# ShortestRoute Project

## Getting started

### Install
This project includes two executables: shortestRoute and cli.
Those files should be created using the following commands:
 
> go install github.com/merjildo/shortestRoute/cli

> go install github.com/merjildo/shortestRoute/shortestRoute

### 1. Run server
You shoud run shortestRoute application in order to enable endpoints

> cd $GOPATH/bin

>./shortestRoute path_to_CSV_file.csv

### 2. Query Interfaces

#### A. Console Interface (Command line - cli):
Open a new console and run the following:

> cd $GOPATH/bin

> ./cli

then you will face the prompt:
"please enter the route:"

#### B. Rest Interfaces
Use Postman, SoapUI or whatever tool you prefer:

##### B.1 Register a new route
In order to reguster a new path, please use:

POST: http://localhost:8080/register 

with the following body:

{

	"From":"BRC",

	"To":"NYC",

	"Weight": 100
}

##### B.2 Consult the shortest route
In the case you want to consult the shortes route
between two pooints, please use:

GET: http://localhost:8080/consult

In the body you have to  set initial and end route:

{	

	"From":"GRU",

	"To":"CDG"
}

### 3. Standalone version

Additionally, this project includes a standalone version. 
This file should be created using the following command:

> go install github.com/merjildo/shortestRoute/standalone

#### 3.1 Run standalone versio

> cd $GOPATH/bin

>./standalone path_to_CSV_file.csv

then you will face the prompt:

> please enter the route: