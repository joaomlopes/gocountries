package gocountries

import (
	"log"
	"testing"
)

func TestAllCountries(t *testing.T) {
	countries, err := AllCountries("")
	if err != nil {
		t.Errorf("Got unexpected error for requested all countries: %v", err)
		return
	}

	if len(countries) < 1 {
		t.Errorf("Unexpected: couldn't find any countries")
	}

	log.Printf("Fetched: %v", (countries)[0])
}

func TestAllCountriesFiltered(t *testing.T) {
	countries, err := AllCountries("name;capital")
	if err != nil {
		t.Errorf("Got unexpected error for requested all countries: %v", err)
		return
	}

	if len(countries) < 1 {
		t.Errorf("Unexpected: couldn't find any countries")
	}

	log.Printf("Fetched: %v", (countries)[0])
}

func TestCountriesByName(t *testing.T) {
	countries, err := CountriesByName("italy", "")
	if err != nil {
		t.Errorf("Got unexpected error for requested country 'italy': %v", err)
		return
	}

	if len(countries) < 1 {
		t.Errorf("Unexpected: couldn't find any country with name 'italy'")
	}

	if (countries)[0].Capital != "Rome" {
		t.Errorf("Unexpected capital for country 'italy'")
	}

	if (countries)[0].RegionalBlocs[0].Acronym != "EU" {
		t.Errorf("Unexpected regional block acronym for country 'italy'")
	}

	log.Printf("Fetched: %v", (countries)[0])
}

func TestCountriesByNameFiltered(t *testing.T) {
	countries, err := CountriesByName("italy", "name;capital;regionalBlocs")
	if err != nil {
		t.Errorf("Got unexpected error for requested country 'italy': %v", err)
		return
	}

	if len(countries) < 1 {
		t.Errorf("Unexpected: couldn't find any country with name 'italy'")
	}

	if (countries)[0].Capital != "Rome" {
		t.Errorf("Unexpected capital for country 'italy'")
	}

	if (countries)[0].RegionalBlocs[0].Acronym != "EU" {
		t.Errorf("Unexpected regional block acronym for country 'italy'")
	}

	log.Printf("Fetched: %v", (countries)[0])
}

func TestCountriesByCapital(t *testing.T) {
	capital := "tallinn"
	countries, err := CountriesByCapital(capital, "")
	if err != nil {
		t.Errorf("Got unexpected error for CountriesByCapital(\"%s\"): %v", capital, err)
		return
	}

	country := (countries)[0]

	if country.Name != "Estonia" {
		t.Errorf("Got unexpected country: expected '%s', got '%s' instead", "Estonia", country.Name)
	}

	log.Printf("Fetched: %v", country)
}

func TestCountriesByCapitalFiltered(t *testing.T) {
	capital := "tallinn"
	countries, err := CountriesByCapital(capital, "name;capital")
	if err != nil {
		t.Errorf("Got unexpected error for CountriesByCapital(\"%s\"): %v", capital, err)
		return
	}

	country := (countries)[0]

	if country.Name != "Estonia" {
		t.Errorf("Got unexpected country: expected '%s', got '%s' instead", "Estonia", country.Name)
	}

	log.Printf("Fetched: %v", country)
}

func TestCountriesByCode(t *testing.T) {
	code := "pt"
	countries, err := CountriesByCode(code, "")
	if err != nil {
		t.Errorf("Got unexpected error for CountriesByCode(\"%s\"): %v", code, err)
		return
	}

	country := countries

	if country.Name != "Portugal" {
		t.Errorf("Got unexpected country: expected '%s', got '%s' instead", "Portugal", country.Name)
	}

	log.Printf("Fetched: %v", country)
}

func TestCountriesByCodeFiltered(t *testing.T) {
	code := "prt"
	countries, err := CountriesByCode(code, "name;capital")
	if err != nil {
		t.Errorf("Got unexpected error for CountriesByCode(\"%s\"): %v", code, err)
		return
	}

	country := countries

	if country.Name != "Portugal" {
		t.Errorf("Got unexpected country: expected '%s', got '%s' instead", "Portugal", country.Name)
	}

	log.Printf("Fetched: %v", country)
}

func TestCountriesByRegion(t *testing.T) {
	region := "europe"
	countries, err := CountriesByRegion(region, "")
	if err != nil {
		t.Errorf("Got unexpected error for CountriesByRegion(\"%s\"): %v", region, err)
		return
	}

	country := (countries)[0]

	if country.Capital != "Mariehamn" {
		t.Errorf("Got unexpected country: expected '%s', got '%s' instead", "Åland Islands", country.Name)
	}

	log.Printf("Fetched: %v", country)
}

func TestCountriesByRegionFiltered(t *testing.T) {
	region := "europe"
	countries, err := CountriesByRegion(region, "name;capital")
	if err != nil {
		t.Errorf("Got unexpected error for CountriesByRegion(\"%s\"): %v", region, err)
		return
	}

	country := (countries)[0]

	if country.Capital != "Mariehamn" {
		t.Errorf("Got unexpected country: expected '%s', got '%s' instead", "Åland Islands", country.Name)
	}

	log.Printf("Fetched: %v", country)
}
