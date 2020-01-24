/**
 * Geo Radius Checker.
 *
 * This package is composed by a utility function used to check if the differences between the coordinates from point a
 * to point b is inside the given radius using the Spherical Law of Cosines.
 *
 */
package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"math"
)

const (
	slcRadius = 6371e3 // Spherical Law of Cosines's radius constant
)

type Coordinates struct {
	ALatitude  float64 `json:"aLatitude"`
	ALongitude float64 `json:"aLongitude"`
	BLatitude  float64 `json:"bLatitude"`
	BLongitude float64 `json:"bLongitude"`
	Radius     int     `json:"radius"` // Radius in meters
}

type Result struct {
	InsideRadius bool `json:"insideRadius"`
}

// Degrees to radians
func toRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func IsInsideRadius(aLatitude float64, aLongitude float64, bLatitude float64, bLongitude float64, radius int) bool {
	diffLongitudeRadians := toRadians(bLongitude - aLongitude)
	aLatitudeRadians := toRadians(aLatitude)
	bLatitudeRadians := toRadians(bLatitude)
	difference := math.Acos(math.Sin(aLatitudeRadians)*math.Sin(bLatitudeRadians)+math.Cos(aLatitudeRadians)*math.Cos(bLatitudeRadians)*math.Cos(diffLongitudeRadians)) * slcRadius
	return difference <= float64(radius)
}

func HandleRequest(ctx context.Context, coordinates Coordinates) (Result, error) {
	result := Result{InsideRadius: IsInsideRadius(coordinates.ALatitude, coordinates.ALongitude, coordinates.BLatitude, coordinates.BLongitude, coordinates.Radius)}
	return result, nil
}

func main() {
	lambda.Start(HandleRequest)
}
