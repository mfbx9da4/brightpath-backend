# Brightpath

## Demo

### Using the demo
- [Demo can be found here](https://brightpath-fe.web.app)
- Click two points anywhere on the map of London.
- The animation between points is just an animation, the response from the server is instant.
- Limited to London due to memory constraints of the free tier of heroku.
- Your live location will be updated regularly if you are walking or cycling.

[
<img src="https://mir-s3-cdn-cf.behance.net/project_modules/1400_opt_1/3036f386223243.5d932aa627487.png" alt="" data-canonical-src="https://mir-s3-cdn-cf.behance.net/project_modules/1400_opt_1/3036f386223243.5d932aa627487.png" width="420" height="345" />
](https://brightpath-fe.web.app)[
<img src="https://user-images.githubusercontent.com/1690659/66380361-30df4d00-e9af-11e9-8579-b569220ddea1.gif" alt="brightpath-demo" data-canonical-src="https://user-images.githubusercontent.com/1690659/66380361-30df4d00-e9af-11e9-8579-b569220ddea1.gif" width="420" height="345" />
](https://brightpath-fe.web.app)

<p align="center">
<img src="https://user-images.githubusercontent.com/1690659/66389489-66406680-e9c0-11e9-8434-309b1b72a63e.png" alt="brightpath-demo" data-canonical-src="https://user-images.githubusercontent.com/1690659/66389489-66406680-e9c0-11e9-8434-309b1b72a63e.png" height="345" />
</p>

## Mission

To provide the safest well lit walking route between two locations.

## Problem

- Google maps can often provide unsafe walking routes.
- If you are new to a city or travelling you are highly dependent on google maps for navigating safely between locations and you will often be on foot.
- There has been a surge of people vying for a product which provides safe walking paths. Here are some [examples](https://twitter.com/Chlojob_/status/1181279485901099008?s=19) from twitter.

<img src="https://user-images.githubusercontent.com/1690659/66380508-83b90480-e9af-11e9-9106-bb1c6ccc59d3.png" alt="" data-canonical-src="https://user-images.githubusercontent.com/1690659/66380508-83b90480-e9af-11e9-9106-bb1c6ccc59d3.png" width="350" /><img src="https://user-images.githubusercontent.com/1690659/66380581-ac40fe80-e9af-11e9-9d36-9933462a1055.png" alt="" data-canonical-src="https://user-images.githubusercontent.com/1690659/66380581-ac40fe80-e9af-11e9-9d36-9933462a1055.png" width="220" /><img src="https://user-images.githubusercontent.com/1690659/66380608-b95ded80-e9af-11e9-937a-013003f04360.png" alt="" data-canonical-src="https://user-images.githubusercontent.com/1690659/66380608-b95ded80-e9af-11e9-937a-013003f04360.png" width="220" />


## Blogs

These were written as a follow up to the hackathon about our experience.

- [My first Hackathon as a UX/UI Designer, (Aristos)](https://medium.com/@aristos.michaelides/my-first-hackathon-as-a-ux-ui-designer-d7fbd7c34602)
- [Route finder app in Golang, (Quynh)](https://medium.com/@dataqween/route-finder-app-100-days-of-golang-day-10-20-bb5f61e21535)
- [LinkedIn post, (Aristos)](https://www.linkedin.com/posts/aristos-michaelides-0639b593_last-weekend-i-participated-in-the-global-activity-6584769414229557248-rPaP/)



## Presentation 

[Presentation slides can be found here](https://docs.google.com/presentation/d/e/2PACX-1vQI5T6knTUv9CPph-cnrhaS_v2JavzzSXZUsJ3H7ZK7uYJaJCTFg4jkwZA7ZjEFGw/pub?start=false&loop=false&delayms=3000)


[
<img src="https://user-images.githubusercontent.com/1690659/66381704-bb28b080-e9b1-11e9-95aa-8fe8bb799f0f.png" alt="" width="400" />
](https://docs.google.com/presentation/d/e/2PACX-1vQI5T6knTUv9CPph-cnrhaS_v2JavzzSXZUsJ3H7ZK7uYJaJCTFg4jkwZA7ZjEFGw/pub?start=false&loop=false&delayms=3000)


## Solution

Brightpath finds walking routes through the city that are well lit, avoiding dark pathways. Later we could incorporate many other kinds of data such as safe checkpoints trivially. The core technical challenge was wrangling the data and writing our own custom routing algorithm. For this hackathon we used data about whether a street is lit or not from [Open Street Map](<[https://www.openstreetmap.org/#map=10/51.4835/-0.1265](https://www.openstreetmap.org/#map=10/51.4835/-0.1265)>). Later we could incorporate official [UK gov data about street lighting](https://data.gov.uk/search?q=Street+Light) or even incorporate satellite imaging data.

[To test the routing algorithm in the wild, click any two locations on the map of London.](https://brightpath-fe.web.app) Note this is limited to London because of the limitations of heroku free plan. While we did optimize for space complexity there are over 5 million edges in London.

### Our approach

Our MVP has focused on using brightness as an objective measure for how safe one might feel. As a society we all need to be more conscious and critical of apps that ‘think for us’, and the data they use. Other ‘objective’ metrics we considered for the future include shop density, heat maps and traffic. We strive to give the power back to the user and allow them to select the safety metrics which are right for them.


## Source code


### Backend

[Backend github repository](https://github.com/mfbx9da4/brightpath-backend). Built in Go, hosted on heroku. https://brightpath.herokuapp.com

Steps taken to create backend

1. Get a dump of London data from Open Street Map
2. Convert OSM format to geojson.
3. [Clean geojson](https://github.com/mfbx9da4/brightpath-backend/blob/master/geojson-cleaner/geojson-cleaner.go) to remove dark pathways.
4. [Write algorithm](https://github.com/mfbx9da4/brightpath-backend/blob/master/parse-geojson.go#L43) to load geojson into graph representation in memory in Go.
5. Create Go API server.
6. Write algorithm for [finding closest known](https://github.com/mfbx9da4/brightpath-backend/blob/master/graph.go#L136) node to user requested start and end coordinates.
7. Write [A\* routing algorithm](https://github.com/mfbx9da4/brightpath-backend/blob/master/graph.go#L164) to find shortest path between two nodes in optimal time.
8. Write [unit tests](https://github.com/mfbx9da4/brightpath-backend/blob/master/routing_test.go) and test cases for shortest path edge cases.

### Frontend 1

- [Frontend 1 github repository](https://github.com/mfbx9da4/brightpath-frontend)
- [Hosted on firebase](https://brightpath-fe.web.app).
- Written in Vanilla JS. 
- Originally written for testing purposes. 
- Uses mapbox for map. 
- Connects to backend hosted on [heroku](https://brightpath.herokuapp.com/).

### Frontend 2

- [Frontend 2 github repository](https://github.com/river-honer/lightpath)
- Uses vuejs
- Uses mapbox geolocation API to get start and end destination.
- Uses leaflet to render map.
- Connects to backend hosted on [heroku](https://brightpath.herokuapp.com/)


## Design Proposal

We put together some [high fidelity mockups](https://www.behance.net/gallery/86223243/BrightPath-Safe-route-finder) of what the app could look like later on.

<p align="center">
<img src="https://user-images.githubusercontent.com/1690659/66382447-1c04b880-e9b3-11e9-8748-ee4ae7fb345f.gif" alt="" height="450" />
</p>

## How it relates to Accessibility and Community

- Our solution is especially useful for those who are travelling and are not familiar with their new surroundings and therefore highly dependent on navigation apps.

- Our solution makes travel more accessible for everyone. We made sure to inlcude minorities which are often overlooked. Our solution does not discriminate - it is for anyone that does not currently feel safe enough to go out at night (whether they are in a different continent, country, city or even an area they are not used to).  These can be individuals with ‘protected characteristics’ under the Equality Act (such as those that are disabled such as partially visually impaired, those with visible differences (such as ethnic minorities or gender) and invisible differences (such as LGBQT+)), but it can also be for any one else that feels vulnerable.

- Our solution also empowers users to engage more with the local community.  For example, from our research there are a number of travellers who will ‘go out at night’, however they will use taxis to get from A to B because they do not feel safe to walk from the hotel to the venue (and/or back again).  I’m sure we all have stories of having some of the best nights when wander freely around a local area, meeting locals serendipitously.  Our app gives more people the opportunity to have impromptu connections with the local community, in more locations.


## Run Backend

```

go build -v -o bin/main . && ./bin/main

```

## Run Frontend

```

yarn install
yarn start

```
