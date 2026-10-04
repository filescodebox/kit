.PHONY: test vet lint clean help

## test: 全量测试(-race)
test:
	go test -race ./...

vet:
	go vet ./...

## lint: golangci-lint(CI 同款门禁)
lint:
	golangci-lint run ./...

clean:
	rm -rf bin/

help:
	@grep -E '^## ' $(MAKEFILE_LIST) | awk '{print $$2}'
