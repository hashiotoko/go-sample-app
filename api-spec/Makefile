.PHONY: start
start:
	@docker ps --filter "name=api-spec" --filter "status=running" --format "{{.Names}}" | grep -w api-spec > /dev/null || \
	docker run -d -p 8880:8080 \
		-e SWAGGER_JSON=/openapi.yaml \
		-v ./generated/openapi.yaml:/openapi.yaml \
		--name api-spec \
		swaggerapi/swagger-ui
