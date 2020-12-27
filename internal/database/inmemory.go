package database

import (
	"context"
	"geobase/internal/model"
)

type InmemoryDB map[string]model.RecyclingPoint

func New() InmemoryDB {
	db := make(InmemoryDB, 10)
	return db
}

func (db InmemoryDB) Init() {
	point01 := model.MapPoint{Latitude: 55.7324321, Longitude: 37.6632525}
	details01 := "20.99z"
	point02 := model.MapPoint{Latitude: 55.7188514, Longitude: 37.574609}
	details02 := "17.97z"
	urlTemplate := "https://www.google.com/maps/@%f,%f,%s"

	recyclingPoint01 := model.RecyclingPoint{Point: point01, Details: details01, UrlTemplate: urlTemplate}
	recyclingPoint02 := model.RecyclingPoint{Point: point02, Details: details02, UrlTemplate: urlTemplate}

	db["стекло"] = recyclingPoint01
	db["пластик"] = recyclingPoint02
	db["метал"] = recyclingPoint01
	db["одежда"] = recyclingPoint02
	db["иное"] = recyclingPoint01
	db["опасное"] = recyclingPoint02
	db["батарейки"] = recyclingPoint01
	db["лампочки"] = recyclingPoint02
	db["техника"] = recyclingPoint01
	db["тетрапак"] = recyclingPoint02
	db["крышечки"] = recyclingPoint01
	db["шины"] = recyclingPoint02
}

func (db InmemoryDB) GetLocationForWasteType(ctx context.Context, recyclingPointRequest model.RecyclingPointRequest) (*model.RecyclingPoint, error) {
	if point, ok := db[recyclingPointRequest.WasteTypeID]; ok {
		return &point, nil
	}
	return nil, ErrNotFound
}
