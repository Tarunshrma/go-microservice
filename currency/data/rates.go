package data

import (
	"encoding/xml"
	"fmt"
	"github.com/hashicorp/go-hclog"
	"net/http"
	"strconv"
)

type ExchangeRates struct {
	l     hclog.Logger
	rates map[string]float64
}

func NewRate(l hclog.Logger) (*ExchangeRates, error) {
	er := &ExchangeRates{l: l, rates: map[string]float64{}}
	err := er.getRates()
	return er, err
}

func (e *ExchangeRates) getRates() error {
	res, err := http.DefaultClient.Get("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml")

	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Expected error code is 200, but recieved %d", res.StatusCode)
	}

	defer res.Body.Close()

	md := &Cubes{}
	xml.NewDecoder(res.Body).Decode(&md)

	for _, c := range md.CubeData {
		r, err := strconv.ParseFloat(c.Rate, 64)
		if err != nil {
			return err
		}

		e.rates[c.Currency] = r
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