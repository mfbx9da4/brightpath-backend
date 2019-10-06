package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type TestData struct {
	GeoJSON      GeoJSON      `json:"geojson"`
	From         Coordinate   `json:"from"`
	To           Coordinate   `json:"to"`
	Distance     float64      `json:"distance"`
	ShortestPath []Coordinate `json:"shortestPath"`
}

func LoadFile(filename string) TestData {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var testData TestData
	json.Unmarshal(byteValue, &testData)
	defer jsonFile.Close()
	return testData
}

func equal(coords1 []Coordinate, coords2 []Coordinate) bool {
	if len(coords1) != len(coords2) {
		return false
	}
	for i := 0; i < len(coords1); i++ {
		if coords1[i] != coords2[i] {
			return false
		}
	}
	return true
}

func TestCreateNode(t *testing.T) {
	var graph Graph
	graph.Init()
	fmt.Println("node1 := graph.GetOrCreateNode([2]float64{0, 0})")
	node1 := Node{0, 0}
	node2 := Node{0, 0}
	node3 := Node{1, 1}
	graph.AddEdge(node2, &node3)
	graph.AddEdge(node3, &node2)
	fmt.Println("graph.Print()")
	graph.Print()
	if len(graph.edges) > 2 {
		t.Errorf("Created too many nodes Got %v", graph.edges)
	}
	if node1 != node2 {
		t.Errorf("Recreating nodes %v %v are not equal", node1, node2)
	}
	if len(graph.edges[node2]) != 1 {
		t.Errorf("Wrong num edges")
	}
	if len(graph.edges[node3]) != 1 {
		t.Errorf("Wrong num edges")
	}
}

func TestPaths(t *testing.T) {
	filenames := []string{
		"./tests/two_points.json",
		"./tests/straight_line.json",
		"./tests/three_roads.json",
		"./tests/shortest_hops.json",
		"./tests/intersection.json",
		"./tests/three_routes.json",
		"./tests/close_but_useless.json",
		"./tests/test01.json",
	}

	for i := 0; i < len(filenames); i++ {
		fmt.Println("===>", filenames[i])
		var expected = LoadFile(filenames[i])
		var graph = createGraph(expected.GeoJSON)
		var route = graph.CalculatePath(expected.From, expected.To)
		var path = getCoordinates(route.Path)
		if !equal(path, expected.ShortestPath) {
			t.Errorf("Incorrect path. Got: %v Expected: %v -- %s", path, expected.ShortestPath, filenames[i])
		}
		if route.Distance != expected.Distance {
			t.Errorf("Incorrect distance. Got: %f Want: %f -- %s", route.Distance, expected.Distance, filenames[i])
		}
	}

}

// type Vertex [2]float64
// type GraphType struct {
// 	vertices Vertices
// 	edges    Edges
// }
// type Vertices map[Vertex]*Vertex
// type Edges map[*Vertex]BagOfVertices
// type BagOfVertices map[*Vertex]bool

// func (graph *GraphType) Init() {
// 	graph.vertices = make(Vertices)
// 	graph.edges = make(Edges)
// }
// func (graph *GraphType) GetOrCreateVertex(vertex Vertex) *Vertex {
// 	if val, ok := graph.vertices[vertex]; ok {
// 		fmt.Println("Found val")
// 		return val
// 	}
// 	graph.vertices[vertex] = &vertex
// 	graph.edges[&vertex] = make(BagOfVertices)
// 	fmt.Println("Create val")
// 	return &vertex
// }

// func TestEdges(t *testing.T) {
// 	var graph GraphType
// 	graph.Init()
// 	// Create vertex 0 and vertex 1
// 	graph.GetOrCreateVertex(Vertex{0, 0})
// 	graph.GetOrCreateVertex(Vertex{1, 1})

// 	// Create edge from vertex 0 to vertex 1
// 	v0 := graph.GetOrCreateVertex(Vertex{0, 0})
// 	v1 := graph.GetOrCreateVertex(Vertex{1, 1})
// 	graph.edges[v0][v1] = true

// 	// Check edge exist from vertex 0 to vertex 1
// 	v0 = graph.GetOrCreateVertex(Vertex{0, 0})
// 	v1 = graph.GetOrCreateVertex(Vertex{1, 1})
// 	if _, ok := graph.edges[v0][v1]; !ok {
// 		t.Errorf("Edge from %v to %v does not exist", v0, v1)
// 	}
// }

// func TestPointers2(t *testing.T) {
// 	var edges Edges = make(map[Vertex]BagOfVertices)
// 	var vertex1 = Vertex{0, 0}
// 	var _vertex1 = Vertex{0, 0}
// 	pointer1 := edges.GetOrCreateVertex(vertex1)
// 	_pointer1 := edges.GetOrCreateVertex(_vertex1)
// 	if vertex1 != _vertex1 {
// 		// Pass
// 		t.Errorf("Values %v %v are not equal", vertex1, _vertex1)
// 	}
// 	if *pointer1 != *_pointer1 {
// 		// Pass
// 		t.Errorf("Values %v %v are not equal", *pointer1, *_pointer1)
// 	}
// 	if pointer1 != _pointer1 {
// 		// Fail, expect to pass
// 		t.Errorf("Pointers %p %p are not equal", pointer1, _pointer1)
// 	}
// }
