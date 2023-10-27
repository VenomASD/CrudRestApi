package main

type Actor struct {
	ActorId   int    ` gorm:"type:int; primaryKey" json:"actorId,string" `
	FirstName string ` gorm:"type:varchar(20)" 		json:"firstName" `
	LastName  string ` gorm:"type:varchar(20)" 		json:"LastName" `
	TimeStamp string ` gorm:"type:timestamp; default: CURRENT_TIMESTAMP(); autoUpdateTime"   		json:"timeStamp" `
}
