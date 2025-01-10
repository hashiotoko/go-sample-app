.PHONY: dev-services/start
dev-services/start:
	docker compose -f docker-compose.yml up
	@echo "please run DB migration if not already."

.PHONY: dev-services/stop
dev-services/stop:
	docker compose -f docker-compose.yml down
