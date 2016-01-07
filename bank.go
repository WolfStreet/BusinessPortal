package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Bank struct {
	ID       bson.ObjectId `bson:"_id" json:"ID"`
	UserName string        `bson:"UserName" json:"UserName"`
	Name     string        `bson:"Name" json:"Name"`
	Actioner string        `bson:"Actioner" json:"Actioner"` //lender or borrower
	Amount   int           `bson:"Amount" json:"Amount"`
	Rate     int           `bson:"Rate" json:"Rate"`
}

func FetchFullBank() []Bank {
	var bank []Bank
	session, err := mgo.DialWithInfo(MongoURL)
	defer session.Close()
	CheckErr("Session error of FetchFullBank", err)
	err = session.DB("RTS").C("Bank").Find(bson.M{}).All(&bank)
	CheckErr("Find error of FetchFullBank", err)
	return bank
}

func FetchBankbyAction(Actioner string) []Bank {
	var bank []Bank
	session, err := mgo.DialWithInfo(MongoURL)
	defer session.Close()
	CheckErr("Session error of FetchBankByAction", err)
	err = session.DB("RTS").C("Bank").Find(bson.M{"Actioner": Actioner}).All(&bank)
	CheckErr("Find error of FetchBankByAction", err)
	return bank
}

func (bank *Bank) FetchBankByID() {
	session, err := mgo.DialWithInfo(MongoURL)
	defer session.Close()
	CheckErr("Session error of FetchBankByID", err)
	err = session.DB("RTS").C("Bank").Find(bson.M{"_id": bank.ID}).One(&bank)
	CheckErr("Find error of FetchBankByID", err)
}

func (bank *Bank) SaveBank() {
	session, err := mgo.DialWithInfo(MongoURL)
	defer session.Close()
	CheckErr("Session error of SaveBank", err)
	err = session.DB("RTS").C("Bank").Insert(&bank)
	err = session.DB("RTS").C("Bank").Update(bson.M{"_id": bank.ID}, &bank)
	CheckErr("Update error of SaveBank", err)
}

func (bank *Bank) PostBank() OutMessage {
	var player Player
	var Message OutMessage
	session, err := mgo.DialWithInfo(MongoURL)
	defer session.Close()
	CheckErr("Session error of PostBank", err)
	err = session.DB("RTS").C("PlayerData").Find(bson.M{"UserName": bank.UserName}).One(&player)
	CheckErr("Player error of PostBank", err)
	bank.Name = player.Name
	if bank.Actioner == "Lender" {
		if player.Career != CAREER_INDEX_BANKER {
			Message.Check = "false"
			Message.Message = "Player is not the Banker. Only Banker can lend."
			return Message
		} else {
			var resources Resources
			resources.Init()
			resources.AssignResource(bank.Amount, "Money")
			if player.isResourceFull(resources) {
				player.DeductResources(resources)
			} else {
				Message.Check = "false"
				Message.Message = "Not enough money."
				return Message
			}
		}
	}
	Message.Check = "true"
	Message.Message = "Bank Request Posted"
	err = session.DB("RTS").C("Bank").Insert(&bank)
	CheckErr("Insert error of PostBank", err)
	return Message
}

func (bank *Bank) AcceptBank(player Player) OutMessage {
	var Message OutMessage
	return Message
}
