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

type ApiGraph struct {
	graph *search.Graph
}

type ApiData struct {
	data           []search.Route
	fileNameRoutes string
}

type ShortestPathResponse struct {
	ShortestPath string
	Distance     int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage : " + os.Args[0] + " filename")
		os.Exit(1)
	}

	filename := os.Args[1]
	data := load.LoadRoutes(filename)
	graph := search.LoadGraph(data)

	apiGraph := &ApiGraph{graph: graph}
	apiRoute := &ApiData{data: data, fileNameRoutes: filename}

	http.HandleFunc("/register", apiRoute.RegisterRoute)
	http.HandleFunc("/consult", apiGraph.ConsultShortestPath)
	// TODO: externalize port
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func (apiData *ApiData) RegisterRoute(w http.ResponseWriter, r *http.Request) {
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

func (apiGraph *ApiGraph) ConsultShortestPath(w http.ResponseWriter, r *http.Request) {
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
