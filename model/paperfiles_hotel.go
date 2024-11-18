package model

type PaperfliesHotel struct {
	HotelId       string `json:"hotel_id"`
	DestinationId int    `json:"destination_id"`
	HotelName     string `json:"hotel_name"`
	Location      struct {
		Address string `json:"address"`
		Country string `json:"country"`
	} `json:"location"`
	Details   string `json:"details"`
	Amenities struct {
		General []string `json:"general"`
		Room    []string `json:"room"`
	} `json:"amenities"`
	Images struct {
		Rooms []struct {
			Link    string `json:"link"`
			Caption string `json:"caption"`
		} `json:"rooms"`
		Site []struct {
			Link    string `json:"link"`
			Caption string `json:"caption"`
		} `json:"site"`
	} `json:"images"`
	BookingConditions []string `json:"booking_conditions"`
}

func MapPaperfliesToHotels(paperfliesHotels []PaperfliesHotel) []Hotel {
	var hotels []Hotel
	for _, paperfliesHotel := range paperfliesHotels {
		location := Location{
			Address: paperfliesHotel.Location.Address,
			Country: paperfliesHotel.Location.Country,
		}

		roomImages := make([]Image, 0, len(paperfliesHotel.Images.Rooms))
		for _, room := range paperfliesHotel.Images.Rooms {
			roomImages = append(roomImages, Image{
				Link:        room.Link,
				Description: room.Caption,
			})
		}

		siteImages := make([]Image, 0, len(paperfliesHotel.Images.Site))
		for _, site := range paperfliesHotel.Images.Site {
			siteImages = append(siteImages, Image{
				Link:        site.Link,
				Description: site.Caption,
			})
		}

		amenities := Amenity{
			General: paperfliesHotel.Amenities.General,
			Room:    paperfliesHotel.Amenities.Room,
		}

		hotels = append(hotels, Hotel{
			ID:            paperfliesHotel.HotelId,
			DestinationID: paperfliesHotel.DestinationId,
			Name:          paperfliesHotel.HotelName,
			Location:      location,
			Amenities:     amenities,
			Images: struct {
				Rooms     []Image `json:"rooms"`
				Site      []Image `json:"site"`
				Amenities []Image `json:"amenities"`
			}{
				Rooms: roomImages,
				Site:  siteImages,
			},
			Description:       paperfliesHotel.Details,
			BookingConditions: paperfliesHotel.BookingConditions,
		})
	}
	return hotels
}
