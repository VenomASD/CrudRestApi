#!/bin/sh

# start mysql server 
docker kill mysql-docker-container

docker rm --force mysql-docker-container

docker pull mysql/mysql-server:latest

docker run -d -p 3306:3306 --name mysql-docker-container -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=mydb -e MYSQL_USER=root -e MYSQL_PASSWORD=root mysql/mysql-server:latest

#start the service 


# ./CrudRestApi