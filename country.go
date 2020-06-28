package gocountries

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL = "https://restcountries.eu/rest/v2/%s"

// Country struct to decode json responses into struct
type Country struct {
	Name           string              `json:"name"`
	Capital        string              `json:"capital"`
	AltSpellings   []string            `json:"altSpellings"`
	Relevance      string              `json:"relevance"`
	Region         string              `json:"region"`
	Subregion      string              `json:"subregion"`
	Translations   map[string]string   `json:"translations"`
	Population     int64               `json:"population"`
	LatLng         []float64           `json:"latlng"`
	Demonym        string              `json:"demonym"`
	Area           float64             `json:"area"`
	Gini           float64             `json:"gini"`
	Timezones      []string            `json:"timezones"`
	Borders        []string            `json:"borders"`
	NativeName     string              `json:"nativeName"`
	CallingCodes   []string            `json:"callingCodes"`
	TopLevelDomain []string            `json:"topLevelDomain"`
	Alpha2Code     string              `json:"alpha2Code"`
	Alpha3Code     string              `json:"alpha3Code"`
	Currencies     []map[string]string `json:"currencies"`
	Languages      []map[string]string `json:"languages"`
	NumericCode    string              `json:"numericCode"`
	Flag           string              `json:"flag"`
	RegionalBlocs  []RegionalBloc      `json:"regionalBlocs"`
	Cioc           string              `json:"cioc"`
}

// RegionalBloc model
type RegionalBloc struct {
	Acronym       string
	Name          string
	OtherAcronyms []string
	OtherNames    []string
}

func doRestcountriesCall(apiSuffix string) ([]byte, error) {
	url := fmt.Sprintf(baseURL, apiSuffix)
	res, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		e := errors.New(fmt.Sprintf("Unexpected API status code %s", res.Status))
		return []byte{}, e
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}

// AllCountries searches for all countries
func AllCountries(filterResponse string) ([]Country, error) {
	apiSuffix := "all"
	if len(filterResponse) > 0 {
		apiSuffix = fmt.Sprintf("all?fields=%s", filterResponse)
	}

	resData, err := doRestcountriesCall(apiSuffix)

	if err != nil {
		return nil, err
	}
	var c []Country
	err = json.Unmarshal(resData, &c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// CountriesByName searches for countries by their name. It can be the native name or partial name
func CountriesByName(name, filterResponse string) ([]Country, error) {
	apiSuffix := fmt.Sprintf("name/%s", name)
	if len(filterResponse) > 0 {
		apiSuffix = fmt.Sprintf("name/%s?fields=%s", name, filterResponse)
	}
	resData, err := doRestcountriesCall(apiSuffix)

	if err != nil {
		return nil, err
	}
	var c []Country
	err = json.Unmarshal(resData, &c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// CountriesByCapital searches for countries with capital city matching 'name'
func CountriesByCapital(name, filterResponse string) ([]Country, error) {
	apiSuffix := fmt.Sprintf("capital/%s", name)
	if len(filterResponse) > 0 {
		apiSuffix = fmt.Sprintf("capital/%s?fields=%s", name, filterResponse)
	}
	resData, err := doRestcountriesCall(apiSuffix)

	if err != nil {
		return nil, err
	}
	var c []Country
	err = json.Unmarshal(resData, &c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// CountriesByCode searches for countries with ISO 3166-1, 2-letter or 3-letter country code
func CountriesByCode(name, filterResponse string) (Country, error) {
	apiSuffix := fmt.Sprintf("alpha/%s", name)
	if len(filterResponse) > 0 {
		apiSuffix = fmt.Sprintf("alpha/%s?fields=%s", name, filterResponse)
	}
	resData, err := doRestcountriesCall(apiSuffix)

	if err != nil {
		return Country{}, err
	}
	var c Country
	err = json.Unmarshal(resData, &c)
	if err != nil {
		return Country{}, err
	}
	return c, nil
}

// CountriesByRegion searches for countries by region matching 'name'
func CountriesByRegion(name, filterResponse string) ([]Country, error) {
	apiSuffix := fmt.Sprintf("region/%s", name)
	if len(filterResponse) > 0 {
		apiSuffix = fmt.Sprintf("region/%s?fields=%s", name, filterResponse)
	}
	resData, err := doRestcountriesCall(apiSuffix)

	if err != nil {
		return nil, err
	}
	var c []Country
	err = json.Unmarshal(resData, &c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
