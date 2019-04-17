package load

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/merjildo/shortestRoute/search"
)

// Routes  return routes loaded from filename CSV file
func Routes(filename string) ([]*search.Route, error) {
	// Open CSV file
	handle, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("%s file name not found", filename)
	}
	defer handle.Close()

	return readData(handle)
}

func readData(handle io.Reader) ([]*search.Route, error) {
	// Read File into a Variable
	lines, err := csv.NewReader(handle).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("Can not read lines from file")
	}

	var data []*search.Route
	for _, line := range lines {
		w, _ := strconv.Atoi(line[2])
		route := search.Route{
			Start:  line[0],
			End:    line[1],
			Weight: w,
		}
		data = append(data, &route)
	}
	return data, nil
}
