package address_test

import (
	"testing"
	. "tests/address"
)

type addressPair struct {
	inputAddress   string
	expectedReturn string
}

func TestAddressType(t *testing.T) {
	tests := []addressPair{
		{"Baker Street", "This is a street address"},
		{"Avenue de la République", "This is a avenue address"},
		{"Boulevard des Capucines", "This is a boulevard address"},
		{"Rue de la Paix", "This is not a valid address type"},
		{"Champs-Élysées", "This is not a valid address type"},
	}

	for _, test := range tests {
		addressType := AddressType(test.inputAddress)
		if addressType != test.expectedReturn {
			t.Errorf("Using input %s, got: %s, want: %s.", test.inputAddress, addressType, test.expectedReturn)
		}
	}
}
