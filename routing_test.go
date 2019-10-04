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
	node1 := graph.GetOrCreateNode([2]float64{0, 0})
	node2 := graph.GetOrCreateNode([2]float64{0, 0})
	graph.Print()
	if len(graph.nodes) > 1 {
		t.Errorf("Created too many nodes Got %v", graph.nodes)
	}
	if node1 != node2 {
		t.Errorf("Recreating nodes %v %v are not equal", node1, node2)
	}
	// if &node1 != &node2 {
	// 	t.Errorf("Recreating nodes %p %p are not equal", &node1, &node2)
	// }
}

// type Store map[int64]int64

// func (s *Store) GetOrCreateByKey(key int64) *int64 {
// 	store := *s
// 	if val, ok := store[key]; ok {
// 		fmt.Println("Found val")
// 		return &val
// 	}
// 	store[key] = key
// 	fmt.Println("Create val")
// 	return &key
// }

// func TestPointers(t *testing.T) {
// 	var myStore = make(Store, 0)
// 	pval1 := myStore.GetOrCreateByKey(1) // Create a node
// 	pval2 := myStore.GetOrCreateByKey(1) // Get it from the store
// 	if pval1 != pval2 {
// 		// Same val, this passes
// 		t.Errorf("Values %v %v are not equal", pval1, pval2)
// 	}
// 	if &pval1 != &pval2 {
// 		// Different pointer, this fails, expected to pass
// 		t.Errorf("Pointers %p %p are not equal", &pval1, &pval2)
// 	}
// }

// func TestPointers(t *testing.T) {
// 	arr := []int{1, 2, 3}
// 	tmp := make([]int, len(arr))
// 	copy(tmp, arr)
// 	path = append(cur.Path, child)
// 	fmt.Println(tmp)
// 	fmt.Println(arr)
// }

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
			t.Errorf("Incorrect path. Got: %v Expected: %v", path, expected.ShortestPath)
		}
		if route.Distance != expected.Distance {
			t.Errorf("Incorrect distance. Got: %f Want: %f", route.Distance, expected.Distance)
		}
	}

}
