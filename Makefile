test:
	@go test -v
	@go test -v ./guard
	@go test -v ./pred
	@go test -v ./run

coverage:
	@go test -v -covermode=count -coverprofile=coverage-tmp.out
	@cat coverage-tmp.out > coverage.out

	@go test -v -covermode=count -coverprofile=coverage-tmp.out ./guard
	@sed -e '1d' < coverage-tmp.out >> coverage.out

	@go test -v -covermode=count -coverprofile=coverage-tmp.out ./pred
	@sed -e '1d' < coverage-tmp.out >> coverage.out

	@go test -v -covermode=count -coverprofile=coverage-tmp.out ./run
	@sed -e '1d' < coverage-tmp.out >> coverage.out

	@rm coverage-tmp.out
