currentPath = `pwd`
loc=`ls -l`

.PHONY: fmt
fmt:
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

.PHONY: lint
lint:
	./bin/golangci-lint run ./...
	
.PHONY: setup
setup:
	@curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.21.0
	go mod tidy

.PHONY: clean
clean:
	rm -rf ./bin