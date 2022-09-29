# CrudRestApi
This is a simple code in golang to implement CRUD REST API.
Database: 
https://www.freesqldatabase.com/account/
use phpmyadmin to manage db

#How to run code:
1) go build (in case you have made some changes in code) : creates a single exe file by using all the go files.
2) .\CrudRestApi.exe : to run the service
3) use postman app to send Http request to the API

#Additional cmds:
1) go mod download : downloads all packages req by project
2) go mod tidy: cleans the go sum file and updates it to latest package info

#Use Swagger to populate API specification

#JenkinsFile

#Docker commands
1) docker build -t crudrestapi:latest .
   This cmd searches docker file in current directory and generates docker image
2) docker run --name crudrestapi -p 8081:8081 crudrestapi:latest
    Runs your image as a container on port 8081
3) use local docker desktop to view your container and its logs