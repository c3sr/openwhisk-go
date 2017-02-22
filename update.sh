#! /bin/bash

set -eou pipefail

SWAGGER_URL="https://raw.githubusercontent.com/openwhisk/openwhisk/master/core/controller/src/main/resources/whiskswagger.json"

if [[ ! -e swagger ]]; then
echo "Download swagger binary..."
latestv=$(curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/latest | jq -r .tag_name)
curl -o swagger -L'#' https://github.com/go-swagger/go-swagger/releases/download/$latestv/swagger_$(echo `uname`|tr '[:upper:]' '[:lower:]')_amd64
chmod +x swagger
else
echo "swagger already present"
fi

echo "Downloading swagger defintion"
curl -o swagger.json "$SWAGGER_URL"

echo "Patching swagger definition"
patch swagger.json swagger.patch

echo "Validating swagger.json"
./swagger validate swagger.json

echo "Generating client"
./swagger generate client -c swagger_client -m swagger_models swagger.json