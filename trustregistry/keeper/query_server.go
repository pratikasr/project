package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/pratikasr/trustregistry"
)

var _ trustregistry.QueryServer = queryServer{}

// NewQueryServerImpl returns an implementation of the module QueryServer.
func NewQueryServerImpl(k Keeper) trustregistry.QueryServer {
	return queryServer{k}
}

type queryServer struct {
	k Keeper
}

// GetTrustRegistry defines the handler for the Query/GetTrustRegistry RPC method.
func (qs queryServer) GetTrustRegistry(ctx context.Context, req *trustregistry.QueryGetTrustRegistryRequest) (*trustregistry.QueryGetTrustRegistryResponse, error) {
	tr, err := qs.k.TrustRegistry.Get(ctx, req.Did)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "trust registry not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	var versions []trustregistry.GovernanceFrameworkVersion
	var documents []trustregistry.GovernanceFrameworkDocument

	// Implement logic to fetch versions and documents based on req.ActiveGfOnly and req.PreferredLanguage
	// This is a placeholder implementation
	if err := qs.k.GFVersion.Walk(ctx, nil, func(key string, gfv trustregistry.GovernanceFrameworkVersion) (bool, error) {
		if gfv.TrDid == req.Did {
			versions = append(versions, gfv)
		}
		return false, nil
	}); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err := qs.k.GFDocument.Walk(ctx, nil, func(key string, gfd trustregistry.GovernanceFrameworkDocument) (bool, error) {
		// Add logic to filter by active version and preferred language
		documents = append(documents, gfd)
		return false, nil
	}); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &trustregistry.QueryGetTrustRegistryResponse{
		TrustRegistry: &tr,
		Versions:      versions,
		Documents:     documents,
	}, nil
}

// Params defines the handler for the Query/Params RPC method.
func (qs queryServer) Params(ctx context.Context, req *trustregistry.QueryParamsRequest) (*trustregistry.QueryParamsResponse, error) {
	params, err := qs.k.Params.Get(ctx)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return &trustregistry.QueryParamsResponse{Params: trustregistry.Params{}}, nil
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &trustregistry.QueryParamsResponse{Params: params}, nil
}
