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

>./shortestRoute

### 2. Interface Testing

#### A. Using Command line (cli):
in a different console run the following:

> cd $GOPATH/bin

> ./cli

then you will face the prompt:
"please enter the route:"

#### B. Using Rest Interfaces
Use Postman, SoapUi or whatever tool you prefer:

##### B.1 Register a new route
please use:

POST: http://localhost:8080/register 

with the following body:

{

	"From":"BRC",

	"To":"NYC",

	"Weight": 100

}

##### B.2 Consult the shortest route
use:

GET: http://localhost:8080/consult

with the following body:

{	

	"From":"GRU",

	"To":"CDG"

}
