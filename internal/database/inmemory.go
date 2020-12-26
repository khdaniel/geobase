package database

import (
	"context"
	"geobase/internal/model"
)

type InmemoryDB map[string]string

func New() InmemoryDB {
	db := make(InmemoryDB, 10)
	return db
}

func (db InmemoryDB) Init() {
	db["Стекло"] = "https://www.google.com/maps/@55.7324321,37.6632525,20.99z"
	db["Пластик"] = "https://www.google.com/maps/@55.7324321,37.6632525,20.99z"
	db["Метал"] = "https://www.google.com/maps/@55.7188514,37.574609,17.97z"
	db["Одежда"] = "https://www.google.com/maps/@55.7188514,37.574609,17.97z"
	db["Иное"] = "https://www.google.com/maps/@55.7188514,37.574609,17.97z"
	db["Опасное"] = "https://www.google.com/maps/@55.7324321,37.6632525,20.99z"
	db["Батарейки"] = "https://www.google.com/maps/@55.7188514,37.574609,17.97z"
	db["Лампочки"] = "https://www.google.com/maps/@55.7188514,37.574609,17.97z"
	db["Техника"] = "https://www.google.com/maps/@55.7188514,37.574609,17.97z"
	db["ТетраПак"] = "https://www.google.com/maps/@55.7188514,37.574609,17.97z"
	db["Крышечки"] = "https://www.google.com/maps/@55.7188514,37.574609,17.97z"
	db["Шины"] = "https://www.google.com/maps/@55.7188514,37.574609,17.97z"
}

func (db InmemoryDB) GetLocationForWasteType(ctx context.Context, recyclingPointRequest model.RecyclingPointRequest) (*model.RecyclingPointDBEntry, error) {
	if url, ok := db[recyclingPointRequest.WasteTypeID]; ok {
		return &model.RecyclingPointDBEntry{WasteType: recyclingPointRequest.WasteTypeID, Url: url}, nil
	}
	return nil, ErrNotFound
}
