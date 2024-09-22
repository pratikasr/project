package trustregistry

import "cosmossdk.io/errors"

var (
	ErrDuplicateDID = errors.Register(ModuleName, 2, "duplicate DID")
	ErrInvalidDID   = errors.Register(ModuleName, 3, "invalid DID")
	ErrInvalidURL   = errors.Register(ModuleName, 4, "invalid URL")
	ErrInvalidHash  = errors.Register(ModuleName, 5, "invalid hash")
)
