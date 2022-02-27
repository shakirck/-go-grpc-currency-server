package data

import (
	"fmt"
	"testing"

	"github.com/hashicorp/go-hclog"
)

func TestRates(t *testing.T) {
	tr, err := NewRates(hclog.Default())
	if err != nil {
		t.Fatal(err)

	}
	fmt.Printf("new rate %#v", tr.Rates)
}
