#!/bin/bash

echo "Generating external check client for JavaScript."
docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate \
    -i /local/spec/swagger.yaml \
    -g javascript \
    -o /local/js/ \
    --minimal-update 
cp -R ${PWD}/spec/templates/js/Dockerfile ${PWD}/js/
echo "Created external check client for JavaScript at ${PWD}/js."

echo "Generating external check client for Python."
docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate \
    -i /local/spec/swagger.yaml \
    -g python \
    -o /local/python/ \
    --minimal-update 
cp -R ${PWD}/spec/templates/python/Dockerfile ${PWD}/python/
echo "Created external check client for Python at ${PWD}/python."

echo "Generating external check client for Java."
docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate \
    -i /local/spec/swagger.yaml \
    -g java \
    -o /local/java/ \
    --minimal-update 
cp -R ${PWD}/spec/templates/java/Dockerfile ${PWD}/java/
echo "Created external check client for Java at ${PWD}/java."

echo "Generating external check client for Go."
docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate \
    -i /local/spec/swagger.yaml \
    -g go \
    -o /local/go/ \
    --minimal-update 
cp -R ${PWD}/spec/templates/go/Dockerfile ${PWD}/go/
echo "Created external check client for Go at ${PWD}/go."