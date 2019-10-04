package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
		result[i] = nodes[i].Value
	}
	return result
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

	geometry := Geometry{Type: "LineString", Coordinates: coords}
	featureOut := Feature{Type: "Feature", ID: "1234", Properties: Property{}, Geometry: geometry}
	features := make([]Feature, 1)

	features[0] = featureOut

	data := GeoJSON{
		Type:     "FeatureCollection",
		Features: features,
	}

	response := FindPathResponse{
		Data:     data,
		Distance: route.Distance}

	geojsonDataInJSON, _ := json.Marshal(&response)

	enableCors(&w)
	w.Write(geojsonDataInJSON)
}

func mapHandler(w http.ResponseWriter, r *http.Request) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened geojson")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	fmt.Println("Successfully ReadAll geojson")
	defer jsonFile.Close()
	enableCors(&w)
	setJSONHeader(&w)
	fmt.Fprintf(w, string(byteValue))
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
	// router.HandleFunc("/map", mapHandler).Methods("GET")
	router.HandleFunc("/map-data/", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		setJSONHeader(&w)
		http.ServeFile(w, r, filename)
	}).Methods("GET")

	fmt.Println("Listening on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, router))

	// terminate if killed
	// go func() {
	//     if err := server.ListenAndServe(":8080", router); err != nil {
	// 		// handle err
	// 		log.Fatal(err)
	//     }
	// }()

	// // Setting up signal capturing
	// stop := make(chan os.Signal, 1)
	// signal.Notify(stop, os.Interrupt)

	// // Waiting for SIGINT (pkill -2)
	// <-stop

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// if err := server.Shutdown(ctx); err != nil {
	//     // handle err
	// }

}
