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
	@rm -rf docs/api
	@mkdir -p docs
	@curl -sL https://github.com/basecamp/fizzy/archive/refs/heads/main.tar.gz | tar -xz -C docs --strip-components=2 fizzy-main/docs/api
	@echo "API spec synced to docs/api"
