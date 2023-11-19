package data

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"testing"
)

func TestNewRates(t *testing.T) {
	tr, error := NewRate(hclog.Default())
	if error != nil {
		t.Fatal(error)
	}

	fmt.Printf("Rates %#v", tr.rates)
}
