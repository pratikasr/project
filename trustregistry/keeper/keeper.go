package keeper

import (
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	storetypes "cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/pratikasr/trustregistry"
)

type Keeper struct {
	cdc          codec.BinaryCodec
	addressCodec address.Codec

	authority string

	Schema        collections.Schema
	Params        collections.Item[trustregistry.Params]
	TrustRegistry collections.Map[string, trustregistry.TrustRegistry]
	GFVersion     collections.Map[string, trustregistry.GovernanceFrameworkVersion]
	GFDocument    collections.Map[string, trustregistry.GovernanceFrameworkDocument]
}

// NewKeeper creates a new Keeper instance
func NewKeeper(cdc codec.BinaryCodec, addressCodec address.Codec, storeService storetypes.KVStoreService, authority string) Keeper {
	if _, err := addressCodec.StringToBytes(authority); err != nil {
		panic(fmt.Errorf("invalid authority address: %w", err))
	}

	sb := collections.NewSchemaBuilder(storeService)
	k := Keeper{
		cdc:          cdc,
		addressCodec: addressCodec,
		authority:    authority,
		Params:       collections.NewItem(sb, trustregistry.ParamsKey, "params", codec.CollValue[trustregistry.Params](cdc)),
		TrustRegistry: collections.NewMap(sb, trustregistry.TrustRegistryKey, "trust_registry",
			collections.StringKey, codec.CollValue[trustregistry.TrustRegistry](cdc)),
		GFVersion: collections.NewMap(sb, trustregistry.GovernanceFrameworkVersionKey, "gf_version",
			collections.StringKey, codec.CollValue[trustregistry.GovernanceFrameworkVersion](cdc)),
		GFDocument: collections.NewMap(sb, trustregistry.GovernanceFrameworkDocumentKey, "gf_document",
			collections.StringKey, codec.CollValue[trustregistry.GovernanceFrameworkDocument](cdc)),
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}

	k.Schema = schema

	return k
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}
