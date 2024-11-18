package model

import "ascenda.assignment.com/utils"

type AcmeHotel struct {
	Id            string      `json:"Id"`
	DestinationId int         `json:"DestinationId"`
	Name          string      `json:"Name"`
	Latitude      interface{} `json:"Latitude"`
	Longitude     interface{} `json:"Longitude"`
	Address       string      `json:"Address"`
	City          string      `json:"City"`
	Country       string      `json:"Country"`
	PostalCode    string      `json:"PostalCode"`
	Description   string      `json:"Description"`
	Facilities    []string    `json:"Facilities"`
}

func MapAcmeToHotels(acmeHotels []AcmeHotel) []Hotel {
	var hotels []Hotel
	for _, acmeHotel := range acmeHotels {
		var lat, lng float32
		if acmeHotel.Latitude != nil && acmeHotel.Longitude != nil {
			lat = utils.ConvertInterfaceToFloat32(acmeHotel.Latitude)
			lng = utils.ConvertInterfaceToFloat32(acmeHotel.Longitude)
		}

		location := Location{
			Lat:     lat,
			Lng:     lng,
			Address: acmeHotel.Address,
			City:    acmeHotel.City,
			Country: acmeHotel.Country,
		}
		hotels = append(hotels, Hotel{
			ID:            acmeHotel.Id,
			DestinationID: acmeHotel.DestinationId,
			Name:          acmeHotel.Name,
			Location:      location,
		})
	}
	return hotels
}
