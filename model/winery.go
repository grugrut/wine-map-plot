package model

import (
	"context"
	"google.golang.org/appengine/datastore"
)

type Winery struct {
	Name      string
	NameJa    string
	Latitude  float64
	Longitude float64
}

// AddWinery add new Winery Entity to datastore
func AddWinery(ctx context.Context, winery Winery) (*datastore.Key, error) {
	return datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "Winery", nil), &winery)
}

func FetchAllWinery(ctx context.Context) ([]Winery, error) {
	var wineries []Winery
	_, err := datastore.NewQuery("Winery").GetAll(ctx, wineries)
	return wineries, err
}
