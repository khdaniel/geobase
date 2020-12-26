package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Cities list
type Cities struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// Storage of app
type Storage struct {
	Cities []Cities `json:"cities"`
}

// New creates a new server
func New(datapath string) *Storage {
	dir, _ := os.Getwd()
	dir = filepath.Join(dir, datapath)
	jFile, err := os.Open(datapath)
	if err != nil {
		fmt.Println(err)
	}
	defer jFile.Close()

	byteValue, _ := ioutil.ReadAll(jFile)

	var store Storage
	json.Unmarshal(byteValue, &store)

	// fmt.Println(store.Cities)

	return &store
}

// GetCities list
func (st *Storage) GetCities() []Cities {
	return st.Cities
}
