package load

import (
	"strings"
	"testing"
)

func TestEmptyFilename(t *testing.T) {
	routes, _ := Routes("")
	if routes != nil {
		t.Errorf("Expected a nil value for routes")
	}
}

func TestReadData(t *testing.T) {
	t.Run("Test something", func(t *testing.T) {
		if _, err := readData(strings.NewReader("GRU, ORL, 20 \n CDG, ORL, 10")); (err != nil) != false {
			t.Errorf("analyze() error = %v", err)
		}
	})
}

func TestEmptyFile(t *testing.T) {
	t.Run("Test something", func(t *testing.T) {
		if _, err := readData(strings.NewReader("")); (err != nil) != false {
			t.Errorf("readData() error = %v", err)
		}
	})
}
