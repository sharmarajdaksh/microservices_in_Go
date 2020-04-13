package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "dexs",
		Price: 1.00,
		SKU:   "abs-bssa-asb",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
