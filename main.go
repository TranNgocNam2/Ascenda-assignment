package main

import (
	"ascenda.assignment.com/model"
	"ascenda.assignment.com/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var (
	acmeSupplierURL       = "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/acme"
	patagoniaSupplierURL  = "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/patagonia"
	paperfliesSupplierURL = "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/paperflies"
)

func fetchHotelsData(url string) ([]model.Hotel, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching data from %s: %w", url, err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var hotels []model.Hotel

	switch url {
	case patagoniaSupplierURL:
		var patagoniaHotels []model.PatagoniaHotel
		err = json.Unmarshal(body, &patagoniaHotels)
		hotels = model.MapPatagoniaToHotels(patagoniaHotels)
		break
	case acmeSupplierURL:
		var acmeHotels []model.AcmeHotel
		err = json.Unmarshal(body, &acmeHotels)
		hotels = model.MapAcmeToHotels(acmeHotels)
		break
	case paperfliesSupplierURL:
		var paperfliesHotels []model.PaperfliesHotel
		err = json.Unmarshal(body, &paperfliesHotels)
		hotels = model.MapPaperfliesToHotels(paperfliesHotels)
		break
	}
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return hotels, nil
}

func mergeHotels(hotels []model.Hotel) map[string]model.Hotel {
	merged := make(map[string]model.Hotel)

	for _, hotel := range hotels {
		if existing, found := merged[hotel.ID]; found {
			existing.Name = utils.ChooseNonEmpty(existing.Name, hotel.Name)
			existing.Description = utils.ChooseNonEmpty(existing.Description, hotel.Description)

			existing.Location.Lat = utils.ChooseNonZero(existing.Location.Lat, hotel.Location.Lat)
			existing.Location.Lng = utils.ChooseNonZero(existing.Location.Lng, hotel.Location.Lng)
			existing.Location.Address = utils.ChooseNonEmpty(existing.Location.Address, hotel.Location.Address)
			existing.Location.City = utils.ChooseNonEmpty(existing.Location.City, hotel.Location.City)
			existing.Location.Country = utils.ChooseNonEmpty(existing.Location.Country, hotel.Location.Country)

			if len(hotel.BookingConditions) > 0 {
				existing.BookingConditions = utils.AppendUniqueStrSlice(existing.BookingConditions, hotel.BookingConditions)
			}

			existing.Amenities.General = utils.AppendUniqueStrSlice(existing.Amenities.General, hotel.Amenities.General)
			existing.Amenities.Room = utils.AppendUniqueStrSlice(existing.Amenities.Room, hotel.Amenities.Room)
			existing.Images.Rooms = append(existing.Images.Rooms, hotel.Images.Rooms...)
			existing.Images.Site = append(existing.Images.Site, hotel.Images.Site...)
			existing.Images.Amenities = append(existing.Images.Amenities, hotel.Images.Amenities...)
			merged[hotel.ID] = existing
		} else {
			merged[hotel.ID] = hotel
		}
	}

	return merged
}

func filterHotels(hotels map[string]model.Hotel, hotelIDs, destinationIDs []string) []model.Hotel {
	var filtered []model.Hotel
	for _, hotel := range hotels {
		destinationMatch := len(destinationIDs) == 0 || utils.ContainsValue(destinationIDs, fmt.Sprintf("%d", hotel.DestinationID))
		hotelMatch := len(hotelIDs) == 0 || utils.ContainsValue(hotelIDs, hotel.ID)

		if destinationMatch && hotelMatch {
			filtered = append(filtered, hotel)
		}
	}
	return filtered
}

func main() {
	hotelIDs := parseInput(os.Args[1])
	destinationIDs := parseInput(os.Args[2])
	patagoniaHotels, err := fetchHotelsData(patagoniaSupplierURL)
	if err != nil {
		fmt.Printf("Error fetching Patagonia hotels: %v\n", err)
		return
	}

	acmeHotels, err := fetchHotelsData(acmeSupplierURL)
	if err != nil {
		fmt.Printf("Error fetching Acme hotels: %v\n", err)
		return
	}

	paperfliesHotels, err := fetchHotelsData(paperfliesSupplierURL)
	if err != nil {
		fmt.Printf("Error fetching Paperflies hotels: %v\n", err)
		return
	}

	var hotels []model.Hotel
	hotels = append(hotels, patagoniaHotels...)
	hotels = append(hotels, acmeHotels...)
	hotels = append(hotels, paperfliesHotels...)

	mergedHotels := mergeHotels(hotels)
	filteredHotels := filterHotels(mergedHotels, hotelIDs, destinationIDs)

	result, err := json.MarshalIndent(filteredHotels, "", " ")
	if err != nil {
		fmt.Printf("Error marshaling filtered hotels: %v\n", err)
		return
	}
	fmt.Println(string(result))
}

func parseInput(input string) []string {
	if input == "none" {
		return nil
	}
	return strings.Split(input, ",")
}
