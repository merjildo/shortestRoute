package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/merjildo/bestRoute/load"
	"github.com/merjildo/bestRoute/search"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage : " + os.Args[0] + " file name")
		os.Exit(1)
	}
	filename := os.Args[1]
	data := load.LoadRoutes(filename)
	graph := search.LoadGraph(data)
	reader := bufio.NewReader(os.Stdin)
	for {
		routeToFind := request(reader)
		shortestPath, shortestDistance := search.Search(graph, routeToFind.Start, routeToFind.End)
		fmt.Println("best route:", strings.Join(shortestPath[:], " - "), " > $", shortestDistance)
	}

}

func request(reader *bufio.Reader) *search.Route {
	fmt.Print("please enter the route: ")
	routeString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	routeString = strings.TrimSuffix(routeString, "\n")
	// TODO: return if empty option (enter)
	routeStr := strings.Split(routeString, "-")
	//TODO: check if node exists
	route := search.Route{
		Start:  routeStr[0],
		End:    routeStr[1],
		Weight: 0,
	}
	return &route
}
