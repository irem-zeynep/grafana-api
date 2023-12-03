## help: print this help message
help:
	@echo "Usage:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ":" | sed -e 's/^/  /'

.PHONY: lint
lint:
	golangci-lint run -c .golangci.yml  --fix -v

.PHONY: test
test:
	go test ./... -coverprofile=unit_coverage.out

.PHONY: unit-coverage-html
unit-coverage-html:
	make test
	go tool cover -html=unit_coverage.out -o unit_coverage.html

.PHONY: generate-mocks
generate-mocks:
	go install github.com/vektra/mockery/v2@latest && mockery --all --keeptree

.PHONY: precommit-install
precommit-install:
	pre-commit install

.PHONY: precommit-run
precommit-run:
	pre-commit run --all-files

.PHONY: before-commit
before-commit: lint test
