package data

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"

	"github.com/hashicorp/go-hclog"
)

type ExchangeRates struct {
	log   hclog.Logger
	Rates map[string]float64
}

func NewRates(l hclog.Logger) (*ExchangeRates, error) {

	ero := &ExchangeRates{log: l, Rates: map[string]float64{}}
	err := ero.getRates()
	return ero, err
}

func (e *ExchangeRates) getRates() error {
	resp, err := http.DefaultClient.Get("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml")

	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Expected 200 got %#v ", resp.StatusCode)

	}
	defer resp.Body.Close()
	md := &Cubes{}
	xml.NewDecoder(resp.Body).Decode(md)

	for _, c := range md.CubeData {
		r, err := strconv.ParseFloat(c.Rate, 64)
		if err != nil {
			return err
		}
		e.Rates[c.Currency] = r
	}
	return nil
}

type Cubes struct {
	CubeData []Cube `xml:"Cube>Cube>Cube"`
}
type Cube struct {
	Currency string `xml:"currency,attr"`
	Rate     string `xml:"rate,attr"`
}
