package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"geobase/internal/database"
	"geobase/internal/model"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func (s *Server) getLocForWasteType(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(s.timeout)*time.Second)
	defer func() {
		s.log.Debug().
			Str("package", "server").
			Str("func", "getLocatio ForWasteType").
			Msg("canceling context")
		cancel()
	}()
	vars := mux.Vars(r)

	wasteTypeID := strings.ToLower(vars["type_id"])

	latitudeParam := r.URL.Query().Get("latitude")
	latitude, err := strconv.ParseFloat(latitudeParam, 32)
	if err != nil {
		http.Error(w, "Unable to get latitude due to:"+err.Error(), http.StatusInternalServerError)
		return
	}

	longitudeParam := r.URL.Query().Get("longitude")
	longitude, err := strconv.ParseFloat(longitudeParam, 32)
	if err != nil {
		http.Error(w, "Unable to get longitude due to:"+err.Error(), http.StatusInternalServerError)
		return
	}

	radiusParam := r.URL.Query().Get("radius")
	radius, err := strconv.ParseInt(radiusParam, 10, 32)
	if err != nil {
		http.Error(w, "Unable to get radius due to:"+err.Error(), http.StatusInternalServerError)
		return
	}

	recyclingPointRequest := model.RecyclingPointRequest{
		WasteTypeID: wasteTypeID,
		Point: model.MapPoint{
			Latitude:  float32(latitude),
			Longitude: float32(longitude),
		},
		Radius: int(radius),
	}

	locationForWasteType, err := s.db.GetLocationForWasteType(ctx, recyclingPointRequest)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			http.Error(w, "No recycling point was found for waste type: "+wasteTypeID, http.StatusNotFound)
			return
		}
	}

	result := model.MapReference{Url: fmt.Sprintf(locationForWasteType.UrlTemplate,
		locationForWasteType.Point.Latitude,
		locationForWasteType.Point.Longitude,
		locationForWasteType.Details)}

	err = json.NewEncoder(w).Encode(&result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) getPointForWasteType(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(s.timeout)*time.Second)
	defer func() {
		s.log.Debug().
			Str("package", "server").
			Str("func", "getPointForWasteType").
			Msg("canceling context")
		cancel()
	}()
	vars := mux.Vars(r)

	wasteTypeID := strings.ToLower(vars["type_id"])

	latitudeParam := r.URL.Query().Get("latitude")
	latitude, err := strconv.ParseFloat(latitudeParam, 32)
	if err != nil {
		http.Error(w, "Unable to get latitude due to:"+err.Error(), http.StatusInternalServerError)
		return
	}

	longitudeParam := r.URL.Query().Get("longitude")
	longitude, err := strconv.ParseFloat(longitudeParam, 32)
	if err != nil {
		http.Error(w, "Unable to get longitude due to:"+err.Error(), http.StatusInternalServerError)
		return
	}

	radiusParam := r.URL.Query().Get("radius")
	radius, err := strconv.ParseInt(radiusParam, 10, 32)
	if err != nil {
		http.Error(w, "Unable to get radius due to:"+err.Error(), http.StatusInternalServerError)
		return
	}

	recyclingPointRequest := model.RecyclingPointRequest{
		WasteTypeID: wasteTypeID,
		Point: model.MapPoint{
			Latitude:  float32(latitude),
			Longitude: float32(longitude),
		},
		Radius: int(radius),
	}

	locationForWasteType, err := s.db.GetLocationForWasteType(ctx, recyclingPointRequest)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			http.Error(w, "No recycling point was found for waste type: "+wasteTypeID, http.StatusNotFound)
			return
		}
	}

	err = json.NewEncoder(w).Encode(&locationForWasteType.Point)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
