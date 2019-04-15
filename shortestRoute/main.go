// Rest Server used to find the shortest route
package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/merjildo/shortestRoute/load"
	"github.com/merjildo/shortestRoute/search"
)

type routeStruct struct {
	From   string
	To     string
	Weight int
}

// APIGraph synthesizes the graph
//used to model the search problem
type APIGraph struct {
	graph *search.Graph
}

// APIData synthesizes the info
// used to load the graphs
type APIData struct {
	data           []search.Route
	fileNameRoutes string
}

// ShortestPathResponse is the desired response.
type ShortestPathResponse struct {
	ShortestPath string
	Distance     int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage : " + os.Args[0] + "  CSV filename (routes.csv)")
		os.Exit(1)
	}

	filename := os.Args[1]
	data := load.LoadRoutes(filename)
	graph := search.LoadGraph(data)

	apiGraph := &APIGraph{graph: graph}
	apiRoute := &APIData{data: data, fileNameRoutes: filename}

	http.HandleFunc("/register", apiRoute.RegisterRoute)
	http.HandleFunc("/consult", apiGraph.ConsultShortestPath)
	// TODO: externalize port
	log.Fatal(http.ListenAndServe(":8080", nil))

}

// RegisterRoute synthesizes an end point
// used to register a new route and save into csv file
func (apiData *APIData) RegisterRoute(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var route routeStruct
	err := decoder.Decode(&route)
	if err != nil {
		panic(err)
	}

	newRoute := search.Route{
		Start:  route.From,
		End:    route.To,
		Weight: route.Weight}

	apiData.data = append(apiData.data, newRoute)

	f, err := os.OpenFile(apiData.fileNameRoutes, os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	writer := csv.NewWriter(f)
	defer writer.Flush()

	// TODO: Verify if the route already exists

	for _, line := range apiData.data {
		strSlice := []string{line.Start, line.End, strconv.Itoa(line.Weight)}
		err = writer.Write(strSlice)

	}
}

// ConsultShortestPath synthesizes an end point
// used to consult the shortest path
func (apiGraph *APIGraph) ConsultShortestPath(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var route routeStruct
	err := decoder.Decode(&route)
	if err != nil {
		panic(err)
	}
	fmt.Println("Requested:", route)

	shortestPath, shortestDistance := search.Search(apiGraph.graph, route.From, route.To)
	response := ShortestPathResponse{
		ShortestPath: strings.Join(shortestPath[:], " - "),
		Distance:     shortestDistance}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}
