package main

const GetDataQuery = `select * from actors where actor_id=%s`

const UpdateDataQuery = `Update actors set first_name="%s", last_name="%s" where actor_id=%s`

const CreateDataQuery = `Insert into actors (actor_id, first_name, last_name) values (%s,"%s","%s")`

const DeleteDataQuery = `Delete from actors where actor_id=%s`
