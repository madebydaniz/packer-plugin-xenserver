#!/bin/bash

set -e

# Load env variables from .env file
if [ -f .env ]; then
  echo "📥 Loading environment from .env"
	set -a
  source .env
	set +a
else
  echo "⚠️  .env file not found! Exiting..."
  exit 1
fi

# Default flag: always include --clean
GORELEASER_FLAGS="--clean $@"

# Extract API_VERSION using go run
echo "🔍 Getting API version from go run . describe ..."
API_VERSION=$(go run . describe | jq -r '.api_version')
export API_VERSION=$API_VERSION

echo "✅ API_VERSION=$API_VERSION"
echo "🔐 GPG_FINGERPRINT=$GPG_FINGERPRINT"
echo "🔑 GITHUB_TOKEN=************"

# Always clean dist folder before build
echo "🧹 Cleaning dist/ directory..."
rm -rf dist/*

# Run build
echo "🏗️ Building packer-plugin-xenserver..."
go build -o packer-plugin-xenserver


# Run GoReleaser with provided flags
echo "🚀 Running: goreleaser release $GORELEASER_FLAGS"
goreleaser release $GORELEASER_FLAGS
