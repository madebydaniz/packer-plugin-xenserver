#!/bin/bash

# Extract API_VERSION using go run
API_VERSION=$(go run . describe | jq -r '.api_version')

# Export env var for goreleaser
export API_VERSION=$API_VERSION

# Run goreleaser with the env var
goreleaser release --clean
