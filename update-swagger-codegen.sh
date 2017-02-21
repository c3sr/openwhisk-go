#! /bin/bash

set -eou pipefail

SWAGGER_URL="https://raw.githubusercontent.com/openwhisk/openwhisk/master/core/controller/src/main/resources/whiskswagger.json"
SWAGGER_JAR=swagger.jar
if [[ ! -e swagger ]]; then
echo "Download swagger jar..."
wget http://repo1.maven.org/maven2/io/swagger/swagger-codegen-cli/2.2.1/swagger-codegen-cli-2.2.1.jar -O $SWAGGER_JAR
else
echo "swagger already present"
fi

echo "Generating client"
java -jar $SWAGGER_JAR generate -l go -i $SWAGGER_URL -o swagger