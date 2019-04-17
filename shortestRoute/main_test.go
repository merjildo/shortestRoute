// Rest Server used to find the shortest route
package main

import (
	"net/http"
	"testing"

	"github.com/merjildo/shortestRoute/search"
)

func TestAPIData_RegisterRoute(t *testing.T) {
	type fields struct {
		data           []*search.Route
		fileNameRoutes string
		graph          *search.Graph
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiData := &APIData{
				data:           tt.fields.data,
				fileNameRoutes: tt.fields.fileNameRoutes,
				graph:          tt.fields.graph,
			}
			apiData.RegisterRoute(tt.args.w, tt.args.r)
		})
	}
}

func TestAPIData_ConsultShortestPath(t *testing.T) {
	type fields struct {
		data           []*search.Route
		fileNameRoutes string
		graph          *search.Graph
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiData := &APIData{
				data:           tt.fields.data,
				fileNameRoutes: tt.fields.fileNameRoutes,
				graph:          tt.fields.graph,
			}
			apiData.ConsultShortestPath(tt.args.w, tt.args.r)
		})
	}
}
