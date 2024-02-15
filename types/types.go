package types

type Distance struct {
	Value float64 `json:"value"`
	OBUID int     `json:"obuID"`
	Unix  int64   `json:"unix"`
}

type OBUData struct {
	OBUId int     `json:"obuID"`
	Lat   float64 `json:"lat"`
	Lon   float64 `json:"lon"`
}
