vet:
	go vet ./...

test:
	go test -race -count=1 ./...

check: vet test-race