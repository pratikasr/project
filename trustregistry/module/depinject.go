package module

import (
	"cosmossdk.io/core/address"
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/core/store"
	"cosmossdk.io/depinject"
	"github.com/pratikasr/trustregistry"

	"github.com/cosmos/cosmos-sdk/codec"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	modulev1 "github.com/pratikasr/trustregistry/api/module/v1"
	"github.com/pratikasr/trustregistry/keeper"
)

var _ appmodule.AppModule = AppModule{}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (am AppModule) IsOnePerModuleType() {}

// IsAppModule implements the appmodule.AppModule interface.
func (am AppModule) IsAppModule() {}

func init() {
	appmodule.Register(
		&modulev1.Module{},
		appmodule.Provide(ProvideModule),
	)
}

type ModuleInputs struct {
	depinject.In

	Cdc          codec.Codec
	StoreService store.KVStoreService
	AddressCodec address.Codec

	Config     *modulev1.Module
	BankKeeper trustregistry.BankKeeper
}

type ModuleOutputs struct {
	depinject.Out

	Module appmodule.AppModule
	Keeper keeper.Keeper
}

func ProvideModule(in ModuleInputs) ModuleOutputs {
	authority := authtypes.NewModuleAddress("gov")
	if in.Config.Authority != "" {
		authority = authtypes.NewModuleAddressOrBech32Address(in.Config.Authority)
	}

	k := keeper.NewKeeper(
		in.Cdc,
		in.AddressCodec,
		in.StoreService,
		authority.String(),
		in.BankKeeper,
	)
	m := NewAppModule(in.Cdc, k)

	return ModuleOutputs{Module: m, Keeper: k}
}
