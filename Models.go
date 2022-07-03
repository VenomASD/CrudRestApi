package main

type Actor struct {
	ActorId   int    ` json:"actorId,string" `
	FirstName string ` json:"firstName" `
	LastName  string ` json:"LastName" `
	TimeStamp string ` json:"timeStamp" `
}
