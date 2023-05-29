
.PHONY: help openapi run

help:
	@echo "Use one of the following commands:"
	@echo "- help"
	@echo "- openapi"
	@echo "- run"

openapi:
	@go run github.com/swaggo/swag/cmd/swag@latest init

run:
	@go run .
