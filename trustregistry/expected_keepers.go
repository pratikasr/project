package trustregistry

import (
	context "context"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BankKeeper sends tokens across modules and is able to get account balances.
type BankKeeper interface {
	GetBalance(ctx context.Context, addr sdk.AccAddress, denom string) sdk.Coin
}
