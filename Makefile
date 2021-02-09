test:
	go test -count=1 -race -cover -v $(shell go list ./... | grep -v -e /vendor/)

lint:
	golint -set_exit_status `go list ./...`

tidy:
	go mod tidy