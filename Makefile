all:
	@echo "'make test' to run tests"
	@echo "'make race' to run tests with -race flag"
	@echo "'make coverage' to run coverage tests"

test:
	@go test -v ./...

race:
	@go test -v -race ./...

coverage:
	@go test -v -covermode=count -coverprofile=coverage-tmp.out
	@cat coverage-tmp.out > coverage.out

	@go test -v -covermode=count -coverprofile=coverage-tmp.out ./errors
	@sed -e '1d' < coverage-tmp.out >> coverage.out

	@go test -v -covermode=count -coverprofile=coverage-tmp.out ./guard
	@sed -e '1d' < coverage-tmp.out >> coverage.out

	@go test -v -covermode=count -coverprofile=coverage-tmp.out ./pred
	@sed -e '1d' < coverage-tmp.out >> coverage.out

	@go test -v -covermode=count -coverprofile=coverage-tmp.out ./run
	@sed -e '1d' < coverage-tmp.out >> coverage.out

	@rm coverage-tmp.out
