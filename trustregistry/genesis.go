package trustregistry

import (
	"fmt"
)

// NewGenesisState creates a new genesis state with default values.
func NewGenesisState() *GenesisState {
	return &GenesisState{
		Params:          DefaultParams(),
		TrustRegistries: []TrustRegistry{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
func (gs GenesisState) Validate() error {
	if err := gs.Params.Validate(); err != nil {
		return err
	}

	uniqueDIDs := make(map[string]bool)
	for _, tr := range gs.TrustRegistries {
		if _, ok := uniqueDIDs[tr.Did]; ok {
			return fmt.Errorf("duplicate DID found: %s", tr.Did)
		}
		uniqueDIDs[tr.Did] = true

		// Basic validation for TrustRegistry fields
		if tr.Did == "" {
			return fmt.Errorf("empty DID in TrustRegistry")
		}
		if tr.Controller == "" {
			return fmt.Errorf("empty Controller in TrustRegistry")
		}
		// Add more field validations as needed
	}

	return nil
}
