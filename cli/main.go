package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/merjildo/bestRoute/search"
)

type Consult struct {
	From string
	To   string
}

type Response struct {
	ShortestPath string
	Distance     int
}

func main() {
	var apiURL string
	var host string
	if len(os.Args) > 2 {
		if os.Args[1] == "--host" {
			host = os.Args[2]
		}
	} else {
		// TODO: Externalize strings
		host = "http://localhost:8080"
	}

	apiURL = host + "/consult"
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("please enter the route: ")
		routeToFind := processRequest(reader)
		consult := Consult{
			From: routeToFind.Start,
			To:   routeToFind.End,
		}
		response := sendRequest(consult, apiURL)
		fmt.Println("best route:", response.ShortestPath, " > $", response.Distance)
	}
}

func sendRequest(consult Consult, apiUrl string) *Response {
	codedConsult, _ := json.Marshal(consult)
	req, err := http.NewRequest(http.MethodGet, apiUrl, bytes.NewBuffer(codedConsult))
	if err != nil {
		panic(err)
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	response := &Response{}
	json.Unmarshal(body, response)
	return response
}

func processRequest(reader *bufio.Reader) *search.Route {
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
