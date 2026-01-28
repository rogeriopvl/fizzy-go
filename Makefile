.PHONY: help run test sync-api-spec

help:
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@grep -E '^[a-zA-Z_-]+:' Makefile | sed 's/:.*//g' | sed 's/^/  /'

run:
	go run .

test:
	gotestsum -- -v ./...

sync-api-spec:
	@mkdir -p docs
	@curl -s https://raw.githubusercontent.com/basecamp/fizzy/main/docs/API.md -o docs/API.md
	@echo "API spec synced to docs/API.md"
