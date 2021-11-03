all: lint test

.PHONY:
lint: lint-deps
	staticcheck ./...

.PHONY:
test:
	go test -v ./...

.PHONY: lint-deps
lint-deps: dep-staticcheck

.PHONY: dep-staticcheck
dep-staticcheck:
	@command -v staticcheck >/dev/null  2>&1 || (echo "missing staticcheck"; go get honnef.co/go/tools/cmd/staticcheck)
