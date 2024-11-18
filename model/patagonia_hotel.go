package model

type PatagoniaHotel struct {
	Id          string   `json:"id"`
	Destination int      `json:"destination"`
	Name        string   `json:"name"`
	Lat         float32  `json:"lat"`
	Lng         float32  `json:"lng"`
	Address     string   `json:"address"`
	Info        *string  `json:"info"`
	Amenities   []string `json:"amenities"`
	Images      struct {
		Rooms []struct {
			Url         string `json:"url"`
			Description string `json:"description"`
		} `json:"rooms"`
		Amenities []struct {
			Url         string `json:"url"`
			Description string `json:"description"`
		} `json:"amenities"`
	} `json:"images"`
}

func MapPatagoniaToHotels(patagoniaHotels []PatagoniaHotel) []Hotel {
	var hotels []Hotel
	for _, patagoniaHotel := range patagoniaHotels {
		roomImages := make([]Image, 0, len(patagoniaHotel.Images.Rooms))
		for _, room := range patagoniaHotel.Images.Rooms {
			roomImages = append(roomImages, Image{
				Link:        room.Url,
				Description: room.Description,
			})
		}

		amenitiesImages := make([]Image, 0, len(patagoniaHotel.Images.Amenities))
		for _, amenity := range patagoniaHotel.Images.Amenities {
			amenitiesImages = append(amenitiesImages, Image{
				Link:        amenity.Url,
				Description: amenity.Description,
			})
		}

		location := Location{
			Lat:     patagoniaHotel.Lat,
			Lng:     patagoniaHotel.Lng,
			Address: patagoniaHotel.Address,
		}

		hotels = append(hotels, Hotel{
			ID:            patagoniaHotel.Id,
			DestinationID: patagoniaHotel.Destination,
			Name:          patagoniaHotel.Name,
			Location:      location,
			Images: struct {
				Rooms     []Image `json:"rooms"`
				Site      []Image `json:"site"`
				Amenities []Image `json:"amenities"`
			}{
				Rooms:     roomImages,
				Amenities: amenitiesImages,
			},
		})
	}
	return hotels
}
