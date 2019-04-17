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

// APIData synthesizes the info
// used to load the graphs
type APIData struct {
	data           []*search.Route
	fileNameRoutes string
	graph          *search.Graph
}

// ShortestPathResponse is the desired response.
type ShortestPathResponse struct {
	ShortestPath string
	Distance     int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage : " + os.Args[0] + " filename.csv")
		os.Exit(1)
	}
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	filename := os.Args[1]
	data, _ := load.Routes(filename)
	graph := search.LoadGraph(data)

	apiRoute := &APIData{
		data:           data,
		fileNameRoutes: filename,
		graph:          graph,
	}

	http.HandleFunc("/register", apiRoute.RegisterRoute)
	http.HandleFunc("/consult", apiRoute.ConsultShortestPath)
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

	f, err := os.OpenFile(apiData.fileNameRoutes, os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	writer := csv.NewWriter(f)
	defer writer.Flush()

	newRoute := &search.Route{
		Start:  route.From,
		End:    route.To,
		Weight: route.Weight,
	}

	checkDuplicate := false
	for _, node := range apiData.data {
		if node.Start == route.From && node.End == route.To {
			log.Println(node.Start + "->" +
				node.End +
				" will be updated to " + strconv.Itoa(route.Weight))
			node.Weight = route.Weight
			checkDuplicate = true
		}
	}
	if !checkDuplicate {
		log.Println(newRoute.Start + "->" + newRoute.End + " will be added.")
		apiData.data = append(apiData.data, newRoute)
	}

	for _, line := range apiData.data {
		strSlice := []string{line.Start, line.End, strconv.Itoa(line.Weight)}
		err = writer.Write(strSlice)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(apiData.data)
}

// ConsultShortestPath synthesizes an end point
// used to consult the shortest path
func (apiData *APIData) ConsultShortestPath(w http.ResponseWriter, r *http.Request) {
	log.Println("Connected from ", r.Host)
	decoder := json.NewDecoder(r.Body)
	var route routeStruct
	err := decoder.Decode(&route)
	if err != nil {
		panic(err)
	}
	log.Println("Payload ", route)

	var response ShortestPathResponse

	if val, msg := checkValues(*apiData, &route); val == true {
		shortestPath, shortestDistance := search.Search(
			apiData.graph,
			strings.ToUpper(route.From),
			strings.ToUpper(route.To))
		response = ShortestPathResponse{
			ShortestPath: strings.Join(shortestPath[:], " - "),
			Distance:     shortestDistance}
	} else {
		response.ShortestPath = msg
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func checkValues(apiData APIData, route *routeStruct) (bool, string) {
	route.From = strings.Trim(route.From, " ")
	route.To = strings.Trim(route.To, " ")
	route.From = strings.ToUpper(route.From)
	route.To = strings.ToUpper(route.To)

	graph := apiData.graph
	var fromOk = false
	var toOk = false
	var msg string
	for _, node := range graph.GetVertices() {
		if route.From == node {
			fromOk = true
		}
		if route.To == node {
			toOk = true
		}
	}

	if !fromOk || !toOk {
		msg = "Error: " + route.To + " or " + route.From + " is wrong"
		log.Println(msg)
	}

	return (fromOk && toOk), msg
}
