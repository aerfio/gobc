.PHONY: fmt
fmt:
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

.PHONY: golang-lint
golang-lint:
	./bin/golangci-lint run ./...
	
.PHONY: setup
setup:
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s v1.21.0
	go mod tidy
