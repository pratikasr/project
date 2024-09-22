package trustregistry

import (
	"fmt"
)

// DefaultParams returns default module parameters.
func DefaultParams() Params {
	return Params{
		MaxDidLength:      100,
		MaxLanguageLength: 2,
		MaxUrlLength:      200,
		MaxHashLength:     64,
	}
}

// Validate does the sanity check on the params.
func (p Params) Validate() error {
	if p.MaxDidLength <= 0 {
		return fmt.Errorf("MaxDIDLength must be positive")
	}
	if p.MaxLanguageLength != 2 {
		return fmt.Errorf("MaxLanguageLength must be 2")
	}
	if p.MaxUrlLength <= 0 {
		return fmt.Errorf("MaxURLLength must be positive")
	}
	if p.MaxHashLength <= 0 {
		return fmt.Errorf("MaxHashLength must be positive")
	}
	return nil
}
