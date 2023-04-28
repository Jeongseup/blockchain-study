# run local celestia
docker run --platform linux/amd64 -p 26650:26657 -p 26659:26659 ghcr.io/rollkit/local-celestia-devnet:v0.9.1

# query test
curl -s -X GET http://0.0.0.0:26659/balance | jq

