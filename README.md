# Chain-Minimal Trustregistry Module

## Introduction

This repository contains the implementation of the trustregistry module 
for the chain-minimal blockchain. The trustregistry module allows for the 
creation and management of trust registries on the blockchain.

## Setup

To set up and run the chain-minimal with the trustregistry module:

1. Navigate to the chain-minimal directory:
   ```
   cd chain-minimal
   ```

2. Install the latest binary:
   ```
   make install
   ```

3. If there are any previous binaries of `minid`, remove them and replace 
with the newly created binary.

4. Initialize the chain:
   ```
   make init
   ```

5. Start the chain:
   ```
   minid start
   ```

## Usage

### Creating a Trust Registry

To create a new trust registry, use the following command:

```bash
minid tx trustregistry create-trust-registry \
did:example:123456789abcdefghi \
http://example-aka.com \
en \
https://example.com/governance-framework.pdf \
e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855 \
--from alice \
--chain-id demo \
--gas auto \
--gas-adjustment 1.3 \
--gas-prices 0.1mini \
--keyring-backend test
```

### Querying Transactions

To query a transaction using its hash:

```bash
minid q tx <txhash>
```

Replace `<txhash>` with the actual transaction hash.

### Querying Trust Registries

To query a trust registry:

```bash
minid q trustregistry get-trust-registry did:example:123456789abcdefghi \
--active-gf-only \
--preferred-language en \
--output json
```

This command will return a JSON response with details about the trust 
registry, including:
- Trust registry information
- Versions of the governance framework
- Documents associated with the trust registry

Example response:

```json
{
  "trust_registry": {
    "did": "did:example:123456789abcdefghi",
    "controller": "mini1uchrwk8q68ywuskd02cn6a7r87matuqlzzrvxl",
    "created": "2024-09-23T05:11:23.873977Z",
    "modified": "2024-09-23T05:11:23.873977Z",
    "aka": "http://example-aka.com",
    "active_version": 1,
    "language": "en"
  },
  "versions": [
    {
      "id": "bf0b181f-d92b-4afa-ab22-0c2d7707634e",
      "tr_did": "did:example:123456789abcdefghi",
      "created": "2024-09-23T05:11:23.873977Z",
      "version": 1,
      "active_since": "2024-09-23T05:11:23.873977Z"
    }
  ],
  "documents": [
    {
      "id": "0b0adef0-f66b-4902-b82b-cfe1e1acd7c3",
      "gfv_id": "bf0b181f-d92b-4afa-ab22-0c2d7707634e",
      "created": "2024-09-23T05:11:23.873977Z",
      "language": "en",
      "url": "https://example.com/governance-framework.pdf",
      "hash": 
"e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
    }
  ]
}
```

## Additional Information

For more detailed information about the trustregistry module and its 
functionality, please refer to the module documentation or contact the 
development team.
