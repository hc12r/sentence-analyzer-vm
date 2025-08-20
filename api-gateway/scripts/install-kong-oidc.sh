#!/bin/bash
# Script to install the Kong OIDC plugin

# Exit on error
set -e

# Check if kubectl is installed
if ! command -v kubectl &> /dev/null; then
  echo "Error: kubectl is not installed."
  exit 1
fi

# Create a ConfigMap for the OIDC Lua plugin
echo "Creating ConfigMap for Kong OIDC plugin..."

# Create a temporary directory
TEMP_DIR=$(mktemp -d)

# Download the Kong OIDC plugin
echo "Downloading Kong OIDC plugin..."
curl -s -L -o $TEMP_DIR/kong-oidc.lua https://raw.githubusercontent.com/nokia/kong-oidc/master/kong/plugins/oidc/handler.lua

# Create the ConfigMap
kubectl create configmap kong-oidc-plugin -n kong --from-file=kong-oidc.lua=$TEMP_DIR/kong-oidc.lua

# Clean up
rm -rf $TEMP_DIR

echo "Kong OIDC plugin installed successfully!"
echo "Now you need to configure your OpenID Connect provider."
echo "Copy api-gateway/config/oidc-config.env.example to oidc-config.env and fill in your provider details."
echo "Then run: source oidc-config.env && ./api-gateway/scripts/setup-oidc.sh"