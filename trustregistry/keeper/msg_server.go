package keeper

import (
	"context"
	"errors"
	"fmt"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pratikasr/trustregistry"
)

type msgServer struct {
	k Keeper
}

var _ trustregistry.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the module MsgServer interface.
func NewMsgServerImpl(keeper Keeper) trustregistry.MsgServer {
	return &msgServer{k: keeper}
}

// CreateTrustRegistry defines the handler for the MsgCreateTrustRegistry message.
func (ms msgServer) CreateTrustRegistry(goCtx context.Context, msg *trustregistry.MsgCreateTrustRegistry) (*trustregistry.MsgCreateTrustRegistryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if _, err := ms.k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, fmt.Errorf("invalid creator address: %w", err)
	}

	if _, err := ms.k.TrustRegistry.Get(ctx, msg.Did); err == nil {
		return nil, fmt.Errorf("trust registry with DID %s already exists", msg.Did)
	} else if !errors.Is(err, collections.ErrNotFound) {
		return nil, err
	}

	tr := trustregistry.TrustRegistry{
		Did:           msg.Did,
		Controller:    msg.Creator,
		Created:       ctx.BlockTime(),
		Modified:      ctx.BlockTime(),
		Deposit:       0, // You might want to implement deposit logic
		Aka:           msg.Aka,
		ActiveVersion: 1,
		Language:      msg.Language,
	}

	if err := ms.k.TrustRegistry.Set(ctx, msg.Did, tr); err != nil {
		return nil, err
	}

	// You might want to create initial GovernanceFrameworkVersion and GovernanceFrameworkDocument here

	return &trustregistry.MsgCreateTrustRegistryResponse{}, nil
}
