package google

import (
	"context"
	"fmt"
	"time"

	"github.com/zopitty/twoshops-be/internal/models"
	"googlemaps.github.io/maps"
)

func FetchOutlets(apiKey, shopName string) ([]models.Outlet, error) {
	client, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	centerLocation := &maps.LatLng{
		Lat: 1.352083,
		Lng: 103.819836,
	}

	var allOutlets []models.Outlet
	var nextPageToken string

	for {
		req := &maps.TextSearchRequest{
			Location: centerLocation,
			Query:    shopName + " Singapore",
			Radius:   22000,
		}
		fmt.Println(req)

		if nextPageToken != "" {
			req.PageToken = nextPageToken
			time.Sleep(2 * time.Second)
		}

		res, err := client.TextSearch(context.Background(), req)
		if err != nil {
			return nil, err
		}

		for _, result := range res.Results {
			if result.BusinessStatus != "OPERATIONAL" {
				continue
			}
			outlet := models.Outlet{
				Name:           result.Name,
				Address:        result.FormattedAddress,
				Latitude:       result.Geometry.Location.Lat,
				Longitude:      result.Geometry.Location.Lng,
				BusinessStatus: result.BusinessStatus,
			}
			allOutlets = append(allOutlets, outlet)
		}
		// break

		nextPageToken = res.NextPageToken
		if nextPageToken == "" {
			break
		}

	}
	return allOutlets, nil
}
