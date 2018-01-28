package model

import (
	"context"
	"google.golang.org/appengine/datastore"
)

type Winery struct {
	Name      string
	Name_ja   string
	longitude float32
	latitude  float32
}

// AddWinery add new Winery Entity to datastore
func AddWinery(ctx context.Context, winery Winery) (*datastore.Key, error) {
	return datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "Winery", nil), &winery)
}
