# we-travel-backend

Kiwi.com hackathon WeTravel API

## Download data

    http://download.geofabrik.de/europe/great-britain/england/greater-london.html

Convert OSM to geojson

    npx osmtogeojson file.osm > file.geojson

## Run API

Install

```
go get -u github.com/gorilla/mux
go get -u github.com/cheekybits/genny/generic
```

To build

```
go build -v -o bin/main .
```

To run

```
go run *.go
```
## Stack and tech
* Choose a small area of London using http://bboxfinder.com/#51.543522,-0.103871,51.546898,-0.094215
* Download Openstreetmap data into .osm https://wiki.openstreetmap.org/wiki/Developer_FAQ
* Convert into geojson using https://github.com/tyrasd/osmtogeojson

Tutorial:
[API tutorial](https://medium.com/the-andela-way/build-a-restful-json-api-with-golang-85a83420c9da)

    curl -o london.osm https://api.openstreetmap.org/api/0.6/map\?bbox\=-0.489,51.680,0.486,51.686

    curl -XPOST -H "Content-type: application/json" -d '{"fromLocation": [-0.100573,51.543750],"toLocation": [-0.097091,51.545466]}' 'http://localhost:8080/findpath'
