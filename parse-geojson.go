package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// Property Geojson Property
type Property struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Highway  string `json:"highway"`
	Access   string `json:"access"`
	Lit      string `json:"lit"`
	Sidewalk string `json:"sidewalk"`
}

// Coordinate pair of lat, lng
type Coordinate = [2]float64

// HashCoordinate return coords as string
func HashCoordinate(coords Coordinate) string {
	var hash = strconv.FormatFloat(coords[0], 'f', -1, 64) + "," + strconv.FormatFloat(coords[1], 'f', -1, 64)
	return hash
}

// Geometry Geojson Geometry
type Geometry struct {
	Type        string       `json:"type"`
	Coordinates []Coordinate `json:"coordinates"`
}

// Feature Geojson Feature
type Feature struct {
	Type       string   `json:"type"`
	ID         string   `json:"id"`
	Properties Property `json:"properties"`
	Geometry   Geometry `json:"geometry"`
}

// GeoJSON data structure
type GeoJSON struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}

func createGraph(geojson GeoJSON) Graph {
	var graph Graph
	for i := 0; i < len(geojson.Features); i++ {
		var feature = geojson.Features[i]
		var prev *Node
		for j := 0; j < len(feature.Geometry.Coordinates); j++ {
			var coords = feature.Geometry.Coordinates[j]
			node := graph.GetOrCreateNode(coords)
			if j != 0 {
				graph.AddEdge(&node, prev)
				graph.AddEdge(prev, &node)
			}
			prev = &node
		}
	}
	fmt.Println("geojson Graph created with", len(graph.nodes), "nodes")
	return graph
}

func loadGeoJSON(filename string) Graph {
	// GeoJSON for Greater London
	// from http://download.geofabrik.de/europe/great-britain/england/greater-london.html
	// geoJsonDownloadLink := "https://ucb7e1be7e59700bb615fc052d06.dl.dropboxusercontent.com/cd/0/get/ApeoomlSroMi4LLrd88j2O1YyfZcz-fnOcR-BMu7Ca3F-aclMpnyLmlzJPZtgze6QSfiGh_SZAcCl-TzGSrcNR14iFsaOBl-vs7CsUzWnL6UbsaH7V_CR-apDThjG8fUH78/file?dl=1DownloadLink"
	// resp, err := http.Get(geoJsonDownloadLink)
	// if err != nil {
	// 	// handle error
	// }
	// defer resp.Body.Close()
	// jsonFile := resp.Body
	// GeoJSON for central london around highbury islington
	// jsonFile, err := os.Open("./data/central.geojson")
	// jsonFile, err := os.Open("./data/greater-london-latest.geojson")
	// 2.73GB
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened geojson")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	fmt.Println("Successfully ReadAll geojson")
	var geojson GeoJSON
	json.Unmarshal(byteValue, &geojson)
	fmt.Println("Successfully Unmarshalled geojson with N features", len(geojson.Features))

	defer jsonFile.Close()
	return createGraph(geojson)
}
