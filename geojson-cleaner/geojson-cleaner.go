package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Property struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Highway  string `json:"highway"`
	Access   string `json:"access"`
	Lit      string `json:"lit"`
	Sidewalk string `json:"sidewalk"`
}

type Coordinate = [2]float64

type Geometry struct {
	Type        string       `json:"type"`
	Coordinates []Coordinate `json:"coordinates"`
}

type Feature struct {
	Type       string   `json:"type"`
	ID         string   `json:"id"`
	Properties Property `json:"properties"`
	Geometry   Geometry `json:"geometry"`
}

type GeoJson struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}

func main() {
	fmt.Println("geojson Graph created with")
	jsonFile, err := os.Open("../data/greater-london-latest.geojson")
	// jsonFile, err := os.Open("../data/central.geojson")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened geojson")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	fmt.Println("Successfully ReadAll geojson", len(byteValue))
	var geojson GeoJson
	json.Unmarshal(byteValue, &geojson)

	var output = GeoJson{Type: geojson.Type, Features: make([]Feature, 0)}

	// var x0 = -0.158044
	// var y0 = 51.546690
	// var x1 = -0.072729
	// var y1 = 51.494244

	var x0 = -0.36782289147757297
	var y0 = 51.54263503347468
	var x1 = 0.05026614032905741
	var y1 = 51.41472350391919

	for i := 0; i < len(geojson.Features); i++ {
		isLineString := geojson.Features[i].Geometry.Type == "LineString"
		isHighway := geojson.Features[i].Properties.Highway != ""
		hasSidewalk := geojson.Features[i].Properties.Sidewalk == "" || geojson.Features[i].Properties.Sidewalk != "none"
		isPath := geojson.Features[i].Properties.Highway == "path"
		isValidPath := geojson.Features[i].Properties.Highway == "path" && (geojson.Features[i].Properties.Access == "no" || geojson.Features[i].Properties.Access == "private")
		isNotPathOrIsValidPath := !isPath || isValidPath
		isLit := geojson.Features[i].Properties.Lit != "" || geojson.Features[i].Properties.Lit == "yes"
		shouldInclude := isHighway && isLineString && hasSidewalk && isNotPathOrIsValidPath && isLit
		// shouldInclude := true
		var feature = geojson.Features[i]
		if shouldInclude && len(feature.Geometry.Coordinates) > 0 {
			var coord = feature.Geometry.Coordinates[0]
			// arbitrary central area to reduce total geojson size
			if coord[0] > x0 && coord[0] < x1 && coord[1] < y0 && coord[1] > y1 {
				output.Features = append(output.Features, feature)
			}
		}
	}
	serialized, _ := json.Marshal(&output)
	err2 := ioutil.WriteFile("../data/cleaned.geojson", serialized, 0644)
	fmt.Println("Reduced features from, to", len(geojson.Features), len(output.Features))
	check(err2)
	defer jsonFile.Close()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
