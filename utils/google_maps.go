package utils

import (
	"go-carbon-tracker/constants"

	"context"
	"fmt"
	"os"

	"googlemaps.github.io/maps"
)

func GetDistances(lat1, lng1, lat2, lng2 float64) (int, error) {
	c, err := maps.NewClient(maps.WithAPIKey(os.Getenv("GOOGLE_MAPS_API_KEY")))
	if err != nil {
		return 0, err
	}

	origins := []string{fmt.Sprintf("%f,%f", lat1, lng1)}
	destinations := []string{fmt.Sprintf("%f,%f", lat2, lng2)}

	req := &maps.DistanceMatrixRequest{
		Origins:      origins,
		Destinations: destinations,
	}

	result, err := c.DistanceMatrix(context.Background(), req)
	if err != nil {
		return 0, err
	}

	if len(result.Rows) > 0 && len(result.Rows[0].Elements) > 0 {
		distance := result.Rows[0].Elements[0].Distance.Meters / 1000

		return distance, nil
	}

	return 0, constants.Err_GET_DISTANCE_FAILED
}

func GetCoordinates(address string) (float64, float64, error) {
	c, err := maps.NewClient(maps.WithAPIKey(os.Getenv("GOOGLE_MAPS_API_KEY")))
	if err != nil {
		return 0, 0, err
	}

	req := &maps.GeocodingRequest{
		Address: address,
	}

	result, err := c.Geocode(context.Background(), req)
	if err != nil {
		return 0, 0, err
	}

	if len(result) > 0 {
		lat := result[0].Geometry.Location.Lat
		lng := result[0].Geometry.Location.Lng
		return lat, lng, nil
	}

	return 0, 0, constants.Err_GET_COORDINATES_FAILED
}
