#!/bin/bash

# Create contracts directory if it doesn't exist
mkdir -p contracts

# Generate bindings for ERC20
echo "Generating ERC20 bindings..."
abigen --abi abis/ERC20.json --pkg contracts --type ERC20 --out contracts/erc20.go

# Generate bindings for EndpointV2
echo "Generating EndpointV2 bindings..."
abigen --abi abis/EndpointV2.json --pkg contracts --type EndpointV2 --out contracts/endpointv2.go

# Generate bindings for EndpointV2View
echo "Generating EndpointV2View bindings..."
abigen --abi abis/EndpointV2View.json --pkg contracts --type EndpointV2View --out contracts/endpointv2view.go

# Generate bindings for IOFT
echo "Generating IOFT bindings..."
abigen --abi abis/ioft.json --pkg contracts --type IOFT --out contracts/ioft.go

echo "Bindings generated successfully!"
