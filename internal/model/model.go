package model

type MapPoint struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type RecyclingPoint struct {
	Point       MapPoint
	Details     string
	UrlTemplate string
}

type RecyclingPointRequest struct {
	WasteTypeID string
	Point       MapPoint
	Radius      int
}

type MapReference struct {
	Url string `json:"url"`
}
