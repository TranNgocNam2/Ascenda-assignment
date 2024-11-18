package model

type Hotel struct {
	ID            string   `json:"id"`
	DestinationID int      `json:"destination_id"`
	Name          string   `json:"name"`
	Location      Location `json:"location"`
	Description   string   `json:"description"`
	Amenities     Amenity  `json:"amenities"`
	Images        struct {
		Rooms     []Image `json:"rooms"`
		Site      []Image `json:"site"`
		Amenities []Image `json:"amenities"`
	} `json:"images"`
	BookingConditions []string `json:"booking_conditions"`
}

type Location struct {
	Lat     float32 `json:"lat"`
	Lng     float32 `json:"lng"`
	Address string  `json:"address"`
	City    string  `json:"city"`
	Country string  `json:"country"`
}

type Image struct {
	Link        string `json:"link"`
	Description string `json:"description"`
}

type Amenity struct {
	General []string `json:"general"`
	Room    []string `json:"room"`
}
