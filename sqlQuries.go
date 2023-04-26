package main

const GetDataQuery = `select * from actor where actor_id=%s`

const GetAllDataQuery = `select * from actor`

const UpdateDataQuery = `Update actor set first_name="%s", last_name="%s" where actor_id=%s`

const CreateDataQuery = `Insert into actor (actor_id, first_name, last_name) values (%s,"%s","%s")`

const DeleteDataQuery = `Delete from actor where actor_id=%s`
