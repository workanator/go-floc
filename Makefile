test:
	@go test -v -covermode=count ./flow
	@go test -v -covermode=count ./guard
	@go test -v -covermode=count ./pred
	@go test -v -covermode=count ./run
	@go test -v -covermode=count ./state
	@go test -v -covermode=count ./

test-coveralls:
	@go test -v -covermode=count -coverprofile=coverage-tmp.out ./flow
	@cat coverage-tmp.out > coverage.out
	@go test -v -covermode=count -coverprofile=coverage-tmp.out ./guard
	@sed -e '1d' < coverage-tmp.out >> coverage.out
	@go test -v -covermode=count -coverprofile=coverage-tmp.out ./pred
	@sed -e '1d' < coverage-tmp.out >> coverage.out
	@go test -v -covermode=count -coverprofile=coverage-tmp.out ./run
	@sed -e '1d' < coverage-tmp.out >> coverage.out
	@go test -v -covermode=count -coverprofile=coverage-tmp.out ./state
	@sed -e '1d' < coverage-tmp.out >> coverage.out
	@go test -v -covermode=count -coverprofile=coverage-tmp.out ./
	@sed -e '1d' < coverage-tmp.out >> coverage.out
	@rm coverage-tmp.out
