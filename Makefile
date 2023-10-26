all: build test
build:
	go build
test:
	go test -v
run:
	./CrudRestApi
# implement me!!
# lint:
# 	golangci-lint 