package trustregistry

import "cosmossdk.io/collections"

const ModuleName = "trustregistry"

var (
	ParamsKey                      = collections.NewPrefix(0)
	TrustRegistryKey               = collections.NewPrefix(1)
	GovernanceFrameworkVersionKey  = collections.NewPrefix(2)
	GovernanceFrameworkDocumentKey = collections.NewPrefix(3)
)
