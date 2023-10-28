all: build test
build:
	go build
test:
	go test -v
run:
	./CrudRestApi local
fmt:
	go fmt
# implement me!!
# lint:
# 	golangci-lint 