#!/bin/bash
set -e

readonly service="$1"
readonly output_dir="$2"
readonly package="$3"

mkdir -p $output_dir
mkdir -p src/app/internal/common/client/$service

oapi-codegen -generate types -o "$output_dir/openapi_types.gen.go" -package "$package" "api/openapi/$service.yml"
oapi-codegen -generate server -o "$output_dir/openapi_api.gen.go" -package "$package" "api/openapi/$service.yml"
oapi-codegen -generate types -o "src/app/internal/common/client/$service/openapi_types.gen.go" -package "$service" "api/openapi/$service.yml"
oapi-codegen -generate client -o "src/app/internal/common/client/$service/openapi_client_gen.go" -package "$service" "api/openapi/$service.yml"