#!/bin/bash
export PATH=$PATH:$(go env GOPATH)/bin

# Create contracts directory if it doesn't exist
mkdir -p contracts

# Create subdirectories for each contract type
mkdir -p contracts/erc20
mkdir -p contracts/endpointv2
mkdir -p contracts/endpointv2view
mkdir -p contracts/ioft

# Generate bindings for ERC20
echo "Generating ERC20 bindings..."
abigen --abi abis/ERC20.json --pkg erc20 --type ERC20 --out contracts/erc20/erc20.go

# Generate bindings for EndpointV2
echo "Generating EndpointV2 bindings..."
abigen --abi abis/EndpointV2.json --pkg endpointv2 --type EndpointV2 --out contracts/endpointv2/endpointv2.go

# Generate bindings for EndpointV2View
echo "Generating EndpointV2View bindings..."
abigen --abi abis/EndpointV2View.json --pkg endpointv2view --type EndpointV2View --out contracts/endpointv2view/endpointv2view.go

# Generate bindings for IOFT
# Generate bindings for IOFT
echo "Generating IOFT bindings..."
abigen --abi abis/IOFT.json --pkg ioft --type IOFT --out contracts/ioft/ioft.go

echo "Bindings generated successfully!"
