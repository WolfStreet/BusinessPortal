//package BusinessPortal
package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func SaveUser(user User) {
	session, err := mgo.DialWithInfo(MongoURL)
	defer session.Close()
	CheckErr("Session error of SaveUser", err)
	err = session.DB("RTS").C("User").Insert(&user)
	CheckErr("Insert error of SaveUser", err)
}

type User struct {
	Id       bson.ObjectId `bson:"_id", omitempty`
	UserName string        `bson:"UserName" json:"UserName"`
	Name     string        `bson:"Name" json:"Name"`
	Password string        `bson:"Password" json:"Password"`
}

func FetchUser(UserName string) bool {
	var user User
	session, err := mgo.DialWithInfo(MongoURL)
	defer session.Close()
	CheckErr("Session error of FetchUser", err)
	err = session.DB("RTS").C("User").Find(bson.M{"UserName": UserName}).One(&user)
	CheckErr("Find error of FetchUser", err)
	return false
}

func ValidUser(user User) bool {
	var temp User
	session, err := mgo.DialWithInfo(MongoURL)
	defer session.Close()
	CheckErr("Session error of ValidUser", err)
	err = session.DB("RTS").C("User").Find(bson.M{"UserName": user.UserName, "Password": user.Password}).One(&temp)
	CheckErr("Find error of ValidUser", err)
	if err != nil {
		return false
	}
	return true
}

func FetchPlayerByName(UserName string) Player {
	var player Player
	var user User
	session, err := mgo.DialWithInfo(MongoURL)
	defer session.Close()
	CheckErr("Session error of FetchPlayerByName", err)
	session.DB("RTS").C("User").Find(bson.M{"UserName": UserName}).One(&user)
	err = session.DB("RTS").C("PlayerData").Find(bson.M{"UserName": UserName}).One(&player)
	if err != nil {
		player.InitPlayer(user.UserName, user.Name, "")
		player.Save()
		return player
	}
	return player
}

func FetchPlayerByCareer(career string) []Player {
	var player []Player
	session, err := mgo.DialWithInfo(MongoURL)
	defer session.Close()
	CheckErr("Session error of FetchPlayerByCareer", err)
	session.DB("RTS").C("PlayerData").Find(bson.M{"Career": career}).All(&player)
	return player
}
