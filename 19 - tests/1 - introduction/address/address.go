package address

import "strings"

// AdressType returns the type of the address
func AddressType(address string) string {
	validTypes := []string{"street", "avenue", "boulevard", "road"}
	for _, validType := range validTypes {
		if strings.Contains(strings.ToLower(address), validType) {
			return "This is a " + validType + " address"
		}
	}
	return "This is not a valid address type"
}
