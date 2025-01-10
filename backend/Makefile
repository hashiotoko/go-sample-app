MIGRATION_DB=mysql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}

.PHONY: run
run:
	go run main.go

.PHONY: migrate/status
migrate/status:
	@echo "Checking migration status"
	atlas migrate status \
		--dir "file://database/migrations" \
		-u ${MIGRATION_DB}

.PHONY: migrate/diff
migrate/diff:
	$(call __require_name)
	@echo "Creating migration file"
	atlas migrate diff "$(name)" \
		--dir "file://database/migrations" \
		--to "file://database/schema.hcl" \
		--dev-url "docker://mysql/8/${DB_NAME}"

.PHONY: migrate/apply
migrate/apply:
	@echo "Applying migration to database"
	atlas migrate apply \
		--dir "file://database/migrations" \
		-u ${MIGRATION_DB}

.PHONY: migrate/new
migrate/new:
	$(call __require_name)
	@echo "Creating new migration"
	atlas migrate new "$(name)" \
		--dir "file://database/migrations"

.PHONY: migrate/hash
migrate/hash:
	@echo "Hashing migration files"
	atlas migrate hash \
		--dir "file://database/migrations"

.PHONY: migrate/clean
migrate/clean:
	@echo "Cleaning up"
	atlas schema clean \
		-u ${MIGRATION_DB}

.PHONY: oapi-code-generate
oapi-code-generate:
	@echo "Generating api interfaces"
	oapi-codegen -generate echo-server -package api ../api-spec/openapi.yaml > api/generated/server.gen.go
	oapi-codegen -generate types -package api ../api-spec/openapi.yaml > api/generated/models.gen.go
	oapi-codegen -generate spec -package api ../api-spec/openapi.yaml > api/generated/spec.gen.go
