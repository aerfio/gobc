.PHONY: fmt
fmt:
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

.PHONY: golang-lint
golang-lint:
	./bin/golangci-lint run ./...

	
# Install all the build and lint dependencies
setup:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh
	go mod tidy
.PHONY: setup