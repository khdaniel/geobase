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
	db["Стекло"] = "https://recyclemap.ru/index.php?id=37334"
	db["Пластик"] = "https://recyclemap.ru/index.php?id=37335"
	db["Метал"] = "https://recyclemap.ru/index.php?id=37336"
	db["Одежда"] = "https://recyclemap.ru/index.php?id=38658"
	db["Иное"] = "https://recyclemap.ru/index.php?id=37305"
	db["Опасное"] = "https://recyclemap.ru/index.php?id=37326"
	db["Батарейки"] = "https://recyclemap.ru/index.php?id=37287"
	db["Лампочки"] = "https://recyclemap.ru/index.php?id=37334"
	db["Техника"] = "https://recyclemap.ru/index.php?id=37335"
	db["ТетраПак"] = "https://recyclemap.ru/index.php?id=37336"
	db["Крышечки"] = "https://recyclemap.ru/index.php?id=37305"
	db["Шины"] = "https://recyclemap.ru/index.php?id=38658"
}

func (db InmemoryDB) GetLocationForWasteType(ctx context.Context, recyclingPointRequest model.RecyclingPointRequest) (*model.RecyclingPointDBEntry, error) {
	if url, ok := db[recyclingPointRequest.WasteTypeID]; ok {
		return &model.RecyclingPointDBEntry{WasteType: recyclingPointRequest.WasteTypeID, Url: url}, nil
	}
	return nil, ErrNotFound
}
