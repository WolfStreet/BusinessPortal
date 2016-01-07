package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Market struct {
	ID       bson.ObjectId `bson:"_id" json:"ID"`
	UserName string        `bson:"UserName" json:"UserName"`
	Name     string        `bson:"Name" json:"Name"`
	Actioner string        `bson:"Actioner" json:"Actioner"` //Buyer or Seller
	Resource string        `bson:"Resource" json:"Resource"`
	Amount   int           `bson:"Amount" json:"Amount"`
	Rate     int           `bson:"Rate" json:"Rate"`
}

func (market Market) InitMarket() {
	market.UserName = ""
	market.Actioner = "Buyer"
	market.Resource = "Food"
	market.Amount = 0
	market.Rate = 0
}

func FetchFullMarket() []Market {
	var market []Market
	session, err := mgo.DialWithInfo(MongoURL)
	defer session.Close()
	CheckErr("Session error of FetchFullMarket", err)
	err = session.DB("RTS").C("Market").Find(bson.M{}).All(&market)
	CheckErr("Document error of FetchFullMarket", err)
	return market

}

func FetchMarketByAction(Actioner string) []Market {
	var market []Market
	session, err := mgo.DialWithInfo(MongoURL)
	defer session.Close()
	CheckErr("Session error of MarketFetchByAction", err)
	err = session.DB("RTS").C("Market").Find(bson.M{"Actioner": Actioner}).All(&market)
	CheckErr("Document error of MarketFetchByAction", err)
	return market
}

func (market *Market) FetchMarketByID() {
	session, err := mgo.DialWithInfo(MongoURL)
	defer session.Close()
	CheckErr("Session error of MarketFetchByID", err)
	err = session.DB("RTS").C("Market").Find(bson.M{"_id": market.ID}).One(&market)
	CheckErr("Document error of MarketFetchByID", err)
}

func (market *Market) SaveMarket() {
	session, err := mgo.DialWithInfo(MongoURL)
	defer session.Close()
	CheckErr("Session error of SaveMarket", err)
	err = session.DB("RTS").C("Market").Update(bson.M{"_id": market.ID}, &market)
	CheckErr("Update error of SaveMarket", err)
}

func (market *Market) PostMarket() OutMessage {
	var player Player
	var Message OutMessage
	session, err := mgo.DialWithInfo(MongoURL)
	defer session.Close()
	CheckErr("Session error of PostMarket", err)
	err = session.DB("RTS").C("PlayerData").Find(bson.M{"UserName": market.UserName}).One(&player)
	CheckErr("Player error of PostMarket", err)
	if market.Actioner == "Seller" {
		if player.Career != CAREER_INDEX_DEALER || player.Career != CAREER_INDEX_ENTREPRENEUR {
			Message.Check = "false"
			Message.Message = "Player is not a dealer/entrpreneur. Only dealer/entrepreneur can sell"
			return Message
		} else {
			var resources Resources
			resources.Init()
			resources.AssignResource(market.Amount, market.Resource)
			if player.isResourceFull(resources) {
				player.DeductResources(resources)
			} else {
				Message.Check = "false"
				Message.Message = "Not Enough Resources"
				return Message
			}
		}
	}
	Message.Check = "true"
	Message.Message = "Market Request Posted"

	err = session.DB("RTS").C("Market").Insert(&market)
	CheckErr("Insert error of PostMarket", err)
	return Message
}

func (market *Market) AcceptMarket(player Player) OutMessage {
	var Message OutMessage
	var resources Resources
	resources.Init()
	resources.AssignResource(market.Amount, market.Resource)
	if market.Actioner == "Buyer" {
		if player.isResourceFull(resources) {
			player.DeductResources(resources)
			Message.Check = "true"
			Message.Message = ""
		} else {
			Message.Check = "false"
			Message.Message = "Not Enough Resources"
		}
	} else if market.Actioner == "Seller" {
		resources.AssignResource(int(market.Amount*market.Rate), RESOURCE_INDEX_MONEY)
		if player.isResourceFull(resources) {
			player.DeductResources(resources)
			resources.AssignResource(market.Amount, market.Resource)
			player.RefundResources(resources)
		} else {
			Message.Check = "false"
			Message.Message = "Not Enough" + RESOURCE_INDEX_MONEY
		}
	} else {
		Message.Check = "false"
		Message.Message = "Could find the action. Please report the bug. Error Code: Market.Accept.01"
	}
	return Message
}
