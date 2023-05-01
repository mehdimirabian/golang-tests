package main

import (
	"fmt"
	"io/ioutil"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	"github.com/paulmach/orb/planar"
)

const (
	GEO_FILE = "points.geojson"
)

func main() {

	// Load in our geojson file into a feature collection
	b, _ := ioutil.ReadFile(GEO_FILE)
	featureCollection, _ := geojson.UnmarshalFeatureCollection(b)

	// Pass in the feature collection + a point of Long/Lat
	if isPointInsidePolygon(featureCollection, orb.Point{100.5, 0.5}) {
		fmt.Println("Point 1 is inside a Polygon")
	} else {
		fmt.Println("Point 1 is not found inside Polygon")
	}

	if isPointInsidePolygon(featureCollection, orb.Point{105.5, 2.5}) {
		fmt.Println("Point 2 is inside a Polygon")
	} else {
		fmt.Println("Point 2 is not found inside Polygon")
	}
}

// isPointInsidePolygon runs through the MultiPolygon and Polygons within a
// feature collection and checks if a point (long/lat) lies within it.
func isPointInsidePolygon(fc *geojson.FeatureCollection, point orb.Point) bool {
	for _, feature := range fc.Features {
		// Try on a MultiPolygon to begin
		multiPoly, isMulti := feature.Geometry.(orb.MultiPolygon)
		if isMulti {
			if planar.MultiPolygonContains(multiPoly, point) {
				return true
			}
		} else {
			// Fallback to Polygon
			polygon, isPoly := feature.Geometry.(orb.Polygon)
			if isPoly {
				if planar.PolygonContains(polygon, point) {
					return true
				}
			}
		}
	}
	return false
}
