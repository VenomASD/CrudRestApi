# CrudRestApi
This is a simple code in golang to implement CRUD REST API.
Database: 
1. install docker desktop from website
2. docker pull mysql
3. docker run -e MYSQL_ROOT_PASSWORD=root -d -p 127.0.0.1:3307:3306 mysql
4. go to docker test terminal and create databse there.
    a) create database mydb; 
    b) create table actor (actor_id int , first_name varchar(40), last_name varchar(40), timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP );


#How to run code:
1) go build (in case you have made some changes in code) : creates a single exe file by using all the go files.
2) .\CrudRestApi.exe : to run the service
3) use postman app to send Http request to the API

#Additional cmds:
1) go mod download : downloads all packages req by project
2) go mod tidy: cleans the go sum file and updates it to latest package info

#Use Swagger to populate API specification

#Jenkins via local docker desktop
1) docker run -p 8080:8080 -p 50000:50000 -d -v jenkins_home:/var/jenkins_home jenkins/jenkins:lts

#Docker commands
1) docker build -t crudrestapi:latest .
   This cmd searches docker file in current directory and generates docker image
2) docker run --name crudrestapi -p 8081:8081 crudrestapi:latest
    Runs your image as a container on port 8081
3) use local docker desktop to view your container and its logs

#pulsar via local docker desktop
1) docker run -it -p 6650:6650  -p 3000:3000 --mount source=pulsardata,target=/pulsar/data --mount source=pulsarconf,target=/pulsar/conf apachepulsar/pulsar:2.10.3 bin/pulsar standalone