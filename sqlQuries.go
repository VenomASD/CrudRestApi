package main

const CreateEmployeeTable = `create table actor (actor_id int , first_name varchar(40), last_name varchar(40), timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP );`

const GetDataQuery = `select * from actor where actor_id=%s`

const UpdateDataQuery = `Update actor set first_name="%s", last_name="%s" where actor_id=%s`

const CreateDataQuery = `Insert into actor (actor_id, first_name, last_name) values (%s,"%s","%s")`

const DeleteDataQuery = `Delete from actor where actor_id=%s`
