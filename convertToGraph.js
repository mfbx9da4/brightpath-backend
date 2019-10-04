const Parser = require("geojson-path-finder");
const geojson = require("./london.geo.json");
const pathFinder = new Parser(geojson);
console.log(JSON.stringify(pathFinder._graph, null, 2));
