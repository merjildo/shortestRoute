package search

// Route is the struct used to set the routes info
type Route struct {
	Start  string
	End    string
	Weight int
}

// Search is the function used to search the shortest route
func Search(g *Graph, start string, end string) ([]string, int) {
	g = resetDistances(g)
	Dijkstra(g, g.GetVertex(start))
	target := g.GetVertex(end)
	shortestPath := []string{target.GetID()}
	Shortest(target, &shortestPath)

	return reverseSlice(shortestPath), target.GetDistance()
}

// LoadGraph is the function used to set routes info.
func LoadGraph(data []*Route) *Graph {
	g := NewGraph()

	nodes := make(map[string]string)
	for _, route := range data {
		nodes[route.Start] = route.Start
		nodes[route.End] = route.End
	}

	for _, vertex := range nodes {
		g.AddVertex(vertex)
	}

	for _, route := range data {
		g.AddEdge(route.Start, route.End, route.Weight)
	}
	return g
}

func reverseSlice(orig []string) []string {
	reversed := []string{}
	for i := range orig {
		n := orig[len(orig)-1-i]
		reversed = append(reversed, n)
	}
	return reversed
}
