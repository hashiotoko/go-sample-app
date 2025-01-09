.PHONY: oapi-code-generate
oapi-code-generate:
	@echo "Generating api interfaces"
	oapi-codegen -generate echo-server -package api ../api-spec/openapi.yaml > api/generated/server.gen.go
	oapi-codegen -generate types -package api ../api-spec/openapi.yaml > api/generated/models.gen.go
	oapi-codegen -generate spec -package api ../api-spec/openapi.yaml > api/generated/spec.gen.go
