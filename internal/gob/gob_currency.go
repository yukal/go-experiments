package gob

import (
	"fmt"

	"golang.org/x/text/currency"
)

type CurrencyWrapper struct {
	currency.Unit
}

func (r CurrencyWrapper) GobEncode() ([]byte, error) {
	if r.Unit.String() == "" {
		return nil, nil
	}

	return []byte(r.Unit.String()), nil
}

func (r *CurrencyWrapper) GobDecode(data []byte) error {
	if len(data) == 0 {
		r.Unit = currency.Unit{}
		return nil
	}

	pattern := string(data)

	if currencyUnit, err := currency.ParseISO(pattern); err != nil {
		return fmt.Errorf("unable parse currency '%s': %w", pattern, err)
	} else {
		r.Unit = currencyUnit
	}

	return nil
}
