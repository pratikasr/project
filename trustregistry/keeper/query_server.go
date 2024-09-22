package keeper

import (
	"context"
	"cosmossdk.io/collections"
	"errors"
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

func (qs queryServer) GetTrustRegistry(ctx context.Context, req *trustregistry.QueryGetTrustRegistryRequest) (*trustregistry.QueryGetTrustRegistryResponse, error) {
	if !isValidDID(req.Did) {
		return nil, status.Error(codes.InvalidArgument, "invalid DID syntax")
	}

	tr, err := qs.k.TrustRegistry.Get(ctx, req.Did)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "trust registry not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	var versions []trustregistry.GovernanceFrameworkVersion
	var documents []trustregistry.GovernanceFrameworkDocument

	// Fetch versions
	err = qs.k.GFVersion.Walk(ctx, nil, func(key string, gfv trustregistry.GovernanceFrameworkVersion) (bool, error) {
		if gfv.TrDid == req.Did && (!req.ActiveGfOnly || gfv.Version == tr.ActiveVersion) {
			versions = append(versions, gfv)
		}
		return false, nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Fetch documents
	err = qs.k.GFDocument.Walk(ctx, nil, func(key string, gfd trustregistry.GovernanceFrameworkDocument) (bool, error) {
		for _, v := range versions {
			if gfd.GfvId == v.Id {
				if req.PreferredLanguage == "" || gfd.Language == req.PreferredLanguage {
					documents = append(documents, gfd)
					break
				} else if len(documents) == 0 || documents[len(documents)-1].GfvId != v.Id {
					documents = append(documents, gfd)
				}
			}
		}
		return false, nil
	})
	if err != nil {
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
