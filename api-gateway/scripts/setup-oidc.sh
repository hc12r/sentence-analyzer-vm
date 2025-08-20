#!/bin/bash
# Script to set up OpenID Connect with Kong

# Exit on error
set -e

# Check if required environment variables are set
if [ -z "$OIDC_CLIENT_ID" ] || [ -z "$OIDC_CLIENT_SECRET" ] || [ -z "$OIDC_DISCOVERY_URL" ]; then
  echo "Error: Required environment variables are not set."
  echo "Please set the following environment variables:"
  echo "  - OIDC_CLIENT_ID: The client ID from your OpenID Connect provider"
  echo "  - OIDC_CLIENT_SECRET: The client secret from your OpenID Connect provider"
  echo "  - OIDC_DISCOVERY_URL: The discovery URL for your OpenID Connect provider"
  echo "  - OIDC_INTROSPECTION_URL: (Optional) The introspection endpoint URL for your OpenID Connect provider"
  echo "                            Not required for Auth0 or other providers that use JWT validation"
  exit 1
fi

# Set default empty value for introspection URL if not provided
OIDC_INTROSPECTION_URL=${OIDC_INTROSPECTION_URL:-""}

# Create a temporary file with the environment variables substituted
TEMP_FILE=$(mktemp)
cat ./api-gateway/kubernetes/kong-config.yaml | \
  sed "s|\${OIDC_CLIENT_ID}|$OIDC_CLIENT_ID|g" | \
  sed "s|\${OIDC_CLIENT_SECRET}|$OIDC_CLIENT_SECRET|g" | \
  sed "s|\${OIDC_DISCOVERY_URL}|$OIDC_DISCOVERY_URL|g" | \
  sed "s|\${OIDC_INTROSPECTION_URL}|$OIDC_INTROSPECTION_URL|g" > $TEMP_FILE

# Apply the configuration
kubectl apply -f $TEMP_FILE

# Clean up
rm $TEMP_FILE

echo "OpenID Connect configuration applied successfully!"