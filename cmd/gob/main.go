package main

import (
	"fmt"
	"log"
	"regexp"
	"time"
	"yu/golang/internal/gob"

	"golang.org/x/text/currency"
)

type Settings struct {
	Language          string // e.g. "UA"
	Location          gob.LocationWrapper
	Currency          gob.CurrencyWrapper
	ValidationPattern gob.RegexpWrapper
}

func initSettings() (Settings, error) {
	locationName := "Europe/Kyiv"
	currencyCode := "UAH" // e.g. "UAH", "USD", "EUR"

	location, err := time.LoadLocation(locationName)
	if err != nil {
		return Settings{}, fmt.Errorf("unable load location: %w", err)
	}

	currencyUnit, err := currency.ParseISO(currencyCode)
	if err != nil {
		return Settings{}, fmt.Errorf("unable parse currency ISO: %w", err)
	}

	// symbol := currency.Symbol(currencyUnit)

	return Settings{
		Language: "UA",
		Location: gob.LocationWrapper{
			LocationName: locationName,
			Location:     location,
		},
		Currency: gob.CurrencyWrapper{
			Unit: currencyUnit,
		},
		ValidationPattern: gob.RegexpWrapper{
			Regexp: regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,6}$`),
		},
	}, nil
}

func main() {
	data, err := initSettings()
	if err != nil {
		log.Fatal("encode error:", err)
	}

	SaveGob("settings.bin", data)
	LoadGob[Settings]("settings.bin")

	// gzipped
	SaveGobz("settings.dat", data)
	LoadGobz[Settings]("settings.dat")
}
