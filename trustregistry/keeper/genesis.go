package keeper

import (
	"context"

	"github.com/pratikasr/trustregistry"
)

// InitGenesis initializes the module state from a genesis state.
func (k *Keeper) InitGenesis(ctx context.Context, data *trustregistry.GenesisState) error {
	if err := k.Params.Set(ctx, data.Params); err != nil {
		return err
	}

	for _, tr := range data.TrustRegistries {
		if err := k.TrustRegistry.Set(ctx, tr.Did, tr); err != nil {
			return err
		}
	}

	return nil
}

// ExportGenesis exports the module state to a genesis state.
func (k *Keeper) ExportGenesis(ctx context.Context) (*trustregistry.GenesisState, error) {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	var trustRegistries []trustregistry.TrustRegistry
	if err := k.TrustRegistry.Walk(ctx, nil, func(did string, tr trustregistry.TrustRegistry) (bool, error) {
		trustRegistries = append(trustRegistries, tr)
		return false, nil
	}); err != nil {
		return nil, err
	}

	return &trustregistry.GenesisState{
		Params:          params,
		TrustRegistries: trustRegistries,
	}, nil
}
