all: build test
build:
	go build
test:
	go test -v
run:
	./CrudRestApi.exe
# implement me!!
# lint:
# 	golangci-lint 