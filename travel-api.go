package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// PathRequestBody API Structure
type PathRequestBody struct {
	FromLocation Coordinate `json:"fromLocation"`
	ToLocation   Coordinate `json:"toLocation"`
}

func getCoordinates(nodes []Node) []Coordinate {
	result := make([]Coordinate, len(nodes))
	for i := 0; i < len(nodes); i++ {
		result[i] = [2]float64{nodes[i][0], nodes[i][1]}
	}
	return result
}

//:::    unit = the unit you desire for results                               :::
//:::           where: 'M' is statute miles (default)                         :::
//:::                  'K' is kilometers                                      :::
//:::                  'N' is nautical miles                                  :::
func distance(coord1 Coordinate, coord2 Coordinate, unit ...string) float64 {
	lng1 := coord1[0]
	lat1 := coord1[1]
	lng2 := coord2[0]
	lat2 := coord2[1]
	const PI float64 = 3.141592653589793
	radlat1 := float64(PI * lat1 / 180)
	radlat2 := float64(PI * lat2 / 180)

	theta := float64(lng1 - lng2)
	radtheta := float64(PI * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / PI
	dist = dist * 60 * 1.1515

	if len(unit) > 0 {
		if unit[0] == "K" {
			dist = dist * 1.609344
		} else if unit[0] == "N" {
			dist = dist * 0.8684
		}
	}

	return dist
}

func calcTotalDistance(coords []Coordinate) float64 {
	sum := 0.0
	for i := 0; i < len(coords); i++ {
		if i > 0 {
			sum += distance(coords[i-1], coords[i], "K")
		}
	}
	return sum
}

func setJSONHeader(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// FindPathResponse API Response
type FindPathResponse struct {
	Data     GeoJSON `json:"data"`
	Distance float64 `json:"distance"`
}

func findpathHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	var newRequestBody PathRequestBody
	if err != nil {
		fmt.Fprintf(w, "Bad input")
	}

	json.Unmarshal(reqBody, &newRequestBody)

	route := graph.CalculatePath(newRequestBody.FromLocation, newRequestBody.ToLocation)
	coords := getCoordinates(route.Path)
	distance := calcTotalDistance(coords)

	geometry := Geometry{Type: "LineString", Coordinates: coords}
	featureOut := Feature{Type: "Feature", ID: "1234", Properties: Property{}, Geometry: geometry}
	features := make([]Feature, 1)

	features[0] = featureOut

	data := GeoJSON{
		Type:     "FeatureCollection",
		Features: features,
	}

	response := FindPathResponse{data, distance}

	geojsonDataInJSON, _ := json.Marshal(&response)

	enableCors(&w)
	setJSONHeader(&w)
	w.Write(geojsonDataInJSON)
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Lightpath backend is running")
}

var graph Graph

// var filename = "./data/greater-london-latest.geojson"

// var filename = "./data/central.geojson"
var filename = "./data/cleaned.geojson"

func main() {
	graph = loadGeoJSON(filename)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/findpath", findpathHandler).Methods("POST")
	router.HandleFunc("/map-data/", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		setJSONHeader(&w)
		http.ServeFile(w, r, filename)
	}).Methods("GET")

	fmt.Println("Listening on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
