# Brightpath

Click a start and end point anywhere on the map of London.
[![button_demo](https://user-images.githubusercontent.com/1690659/66333882-c63afc80-e92f-11e9-8768-1fcdd4411359.png)](https://brightpath-fe.web.app)

[
![brightpath](https://mir-s3-cdn-cf.behance.net/project_modules/1400_opt_1/3036f386223243.5d932aa627487.png)
](https://brightpath-fe.web.app)

[
![brightpath-demo](https://user-images.githubusercontent.com/1690659/66332818-807d3480-e92d-11e9-9c07-a35f045ef92e.gif)
](https://brightpath-fe.web.app)

[![button_demo](https://user-images.githubusercontent.com/1690659/66333882-c63afc80-e92f-11e9-8768-1fcdd4411359.png)](https://brightpath-fe.web.app)

## Mission

To provide the safest walking route between two locations.

## Problem

- Google maps can often provide unsafe walking routes.
- If you are new to a city or travelling you are highly dependent on google maps for navigating safely between locations and you will often be on foot.
- There has been a surge of people vying for a product which provides safe walking paths. Here are [some](https://twitter.com) [examples](https://twitter.com) on [twitter](https://twitter.com).

## Solution

Brightpath finds walking routes through the city that are well lit, avoiding dark pathways. Later we could incorporate many other kinds of data such as crime reports and safe checkpoints. The core technical challenge was wrangling the data and writing our own custom routing algorithm. For this hackathon we used data about whether a street is lit or not from [Open Street Map](<[https://www.openstreetmap.org/#map=10/51.4835/-0.1265](https://www.openstreetmap.org/#map=10/51.4835/-0.1265)>). Later we could incorporate official [UK gov data about street lighting](https://data.gov.uk/search?q=Street+Light) or even incorporate satellite imaging data.

[To use the routing in action, click any two locations on the map.](https://brightpath-fe.web.app) Note this is limited to London because of the limitations of heroku free plan. While we did optimize for space complexity there are over 5 million edges in London.

## Technical Solution

[Backend github repository](https://github.com/mfbx9da4/brightpath-backend)
[Frontend 1 github repository](https://github.com/mfbx9da4/brightpath-frontend)
[Frontend 2 github repository](https://github.com/river-honer/lightpath)

### Backend

Built in Go, hosted on heroku. https://brightpath.herokuapp.com

Steps taken to create backend

1. Get a dump of London data from [Open Street Map](http://download.geofabrik.de/europe/great-britain/england/greater-london.html)
2. Convert OSM format to geojson.
3. Clean geojson to remove dark pathways.
4. [Write algorithm](https://github.com/mfbx9da4/brightpath-backend/blob/master/parse-geojson.go#L43) to load geojson into graph representation in memory in Go.
5. Create Go API server.
6. Write algorithm for finding closest node two api source and output destination.
7. Write [A\* routing algorithm](https://github.com/mfbx9da4/brightpath-backend/blob/master/graph.go#L164) to find shortest path between two nodes in optimal time.
8. Write [unit tests](https://github.com/mfbx9da4/brightpath-backend/blob/master/routing_test.go) and test cases for shortest path edge cases.

### Frontend 1

[Hosted on firebase](https://brightpath-fe.web.app), written in Vanilla JS. Originally written for testing purposes. Uses mapbox for map. Connects to backend hosted on [heroku](https://brightpath.herokuapp.com/).

### Frontend 2

Uses vuejs

- Uses mapbox geolocation API to get start and end destination.
- Uses leaflet to render map.

### Hackathon topic you used

Accessibility & Community

### The development tools used to build the project

- Golang (Backend)

- Vue.js (Frontend)

## LightPath Backend Repo

The backend is a Go project who's github is here [https://github.com/cryptoqween/we-travel-backend](https://github.com/cryptoqween/we-travel-backend)

## Design Proposal

https://www.behance.net/gallery/86223243/BrightPath-Safe-route-finder

## Run Backend

```

go build -v -o bin/main . && ./bin/main

```

## Run Frontend

```

yarn install
yarn start

```
