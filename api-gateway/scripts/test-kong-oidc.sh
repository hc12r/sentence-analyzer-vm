#!/bin/bash
# Script to test Kong API Gateway with OpenID Connect

# Exit on error
set -e

# Check if required environment variables are set
if [ -z "$OIDC_CLIENT_ID" ] || [ -z "$OIDC_CLIENT_SECRET" ] || [ -z "$OIDC_TOKEN_URL" ]; then
  echo "Error: Required environment variables are not set."
  echo "Please set the following environment variables:"
  echo "  - OIDC_CLIENT_ID: The client ID from your OpenID Connect provider"
  echo "  - OIDC_CLIENT_SECRET: The client secret from your OpenID Connect provider"
  echo "  - OIDC_TOKEN_URL: The token endpoint URL for your OpenID Connect provider"
  exit 1
fi

# Get the server IP
if [ -z "$SERVER_IP" ]; then
  echo "SERVER_IP environment variable not set. Please provide the server IP address."
  exit 1
fi

echo "Testing Kong API Gateway with OpenID Connect..."
echo "Server IP: $SERVER_IP"

# Step 1: Check if Kong is running
echo -e "\n1. Checking if Kong is running..."
KONG_STATUS=$(curl -s -o /dev/null -w "%{http_code}" http://$SERVER_IP:8001 || echo "Failed")
if [ "$KONG_STATUS" = "200" ]; then
  echo "✅ Kong is running and the Admin API is accessible."
else
  echo "❌ Kong is not running or the Admin API is not accessible. Status code: $KONG_STATUS"
  exit 1
fi

# Step 2: Get an access token from the OpenID Connect provider
echo -e "\n2. Getting an access token from the OpenID Connect provider..."
TOKEN_RESPONSE=$(curl -s -X POST \
  $OIDC_TOKEN_URL \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -d "grant_type=client_credentials&client_id=$OIDC_CLIENT_ID&client_secret=$OIDC_CLIENT_SECRET")

# Extract the access token
ACCESS_TOKEN=$(echo $TOKEN_RESPONSE | grep -o '"access_token":"[^"]*' | cut -d'"' -f4)

if [ -z "$ACCESS_TOKEN" ]; then
  echo "❌ Failed to get an access token. Response: $TOKEN_RESPONSE"
  exit 1
else
  echo "✅ Successfully obtained an access token."
fi

# Step 3: Test the API with the access token
echo -e "\n3. Testing the API with the access token..."
API_RESPONSE=$(curl -s -X POST \
  http://$SERVER_IP:30080/analyze \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $ACCESS_TOKEN" \
  -d '{"sentence": "Hello World"}')

# Check if the API response contains the expected fields
if echo $API_RESPONSE | grep -q "word_count"; then
  echo "✅ API request successful! Response: $API_RESPONSE"
else
  echo "❌ API request failed. Response: $API_RESPONSE"
  exit 1
fi

# Step 4: Test the API without an access token (should fail)
echo -e "\n4. Testing the API without an access token (should fail)..."
UNAUTHORIZED_RESPONSE=$(curl -s -o /dev/null -w "%{http_code}" -X POST \
  http://$SERVER_IP:30080/analyze \
  -H 'Content-Type: application/json' \
  -d '{"sentence": "Hello World"}')

if [ "$UNAUTHORIZED_RESPONSE" = "401" ]; then
  echo "✅ API correctly rejected the request without an access token."
else
  echo "❌ API did not reject the request without an access token. Status code: $UNAUTHORIZED_RESPONSE"
  exit 1
fi

echo -e "\n✅ All tests passed! Kong API Gateway with OpenID Connect is working correctly."