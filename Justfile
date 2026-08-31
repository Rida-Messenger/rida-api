fmt:
    gofmt -w .

fmt-check:
    @test -z "$(gofmt -l .)" || ( \
        echo "Files are not formatted:"; \
        gofmt -l .; \
        exit 1 \
    )

tidy:
    go mod tidy

tidy-check:
    go mod tidy -diff

vet:
    go vet ./...

test: 
    go test ./...

build:
    go build ./...

check-ci: fmt-check tidy-check vet test build