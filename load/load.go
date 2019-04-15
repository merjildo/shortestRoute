package load

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/merjildo/shortestRoute/search"
)

func LoadRoutes(filename string) []search.Route {
	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	var data []search.Route

	for _, line := range lines {
		w, _ := strconv.Atoi(line[2])
		route := search.Route{
			Start:  line[0],
			End:    line[1],
			Weight: w,
		}
		data = append(data, route)
	}
	return data
}
