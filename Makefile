all: build test
build:
	go build
test:
	go test -v
run:
	./CrudRestApi.exe
jenkins:
	docker run -p 8080:8080 -p 50000:50000 -d -v jenkins_home:/var/jenkins_home jenkins/jenkins:lts
# implement me!!
# lint:
# 	golangci-lint 