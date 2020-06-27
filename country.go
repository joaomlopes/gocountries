package gocountries

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL = "https://restcountries.eu/rest/v2/%s"

type Country struct {
	Name           string
	Capital        string
	AltSpellings   []string
	Relevance      string
	Region         string
	Subregion      string
	Translations   map[string]string
	Population     int32
	LatLng         []float32
	Demonym        string
	Area           float32
	Gini           float32
	Timezones      []string
	Borders        []string
	NativeName     string
	CallingCodes   []string
	TopLevelDomain []string
	Alpha2Code     string
	Alpha3Code     string
	Currencies     []map[string]string
	Languages      []map[string]string
	NumericCode    string
	Flag           string
	RegionalBlocs  []RegionalBloc
	Cioc           string
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
