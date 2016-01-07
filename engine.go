package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"os"
	"time"
)

//var MongoURL = "mongodb://wolfstreet:H01204134192h@ds059722.mongolab.com:59722"
var MongoURL = &mgo.DialInfo{
	Addrs:    []string{"127.0.0.1:27017"},
	Database: "RTS",
	FailFast: true,
}

type Clock struct {
	Hour   int
	Min    int
	Second int
}

var MidNight = Clock{0, 0, 0}
var LogFile, _ = os.OpenFile("./log/log.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0600)
var Logger = log.New(LogFile, "\n", log.LstdFlags)

const (
	tic = 1 * time.Second
)

func CheckErr(note string, err error) {
	if err != nil {
		fmt.Println(note)
		fmt.Println(err)
	}
}

func CatchTic() {
	var AllPlayers []Player
	session, err := mgo.DialWithInfo(MongoURL)
	//session, err := mgo.Dial("localhost:27017")
	CheckErr("Catch Tic: ", err)
	defer session.Close()
	for range time.Tick(tic) {
		err = session.DB("RTS").C("PlayerData").Find(bson.M{}).All(&AllPlayers)
		CheckErr("Inside Tic: ", err)
		for _, player := range AllPlayers {

			player.ResStored.Money += player.ResProduction.Money * float32(tic/time.Second) / 3600.0
			player.ResStored.Food += player.ResProduction.Food * float32(tic/time.Second) / 3600.0
			player.ResStored.Wood += player.ResProduction.Wood * float32(tic/time.Second) / 3600.0
			player.ResStored.Metal += player.ResProduction.Metal * float32(tic/time.Second) / 3600.0
			player.ResStored.Cement += player.ResProduction.Cement * float32(tic/time.Second) / 3600.0
			player.ResStored.Clothes += player.ResProduction.Clothes * float32(tic/time.Second) / 3600.0
			player.ResStored.Fuel += player.ResProduction.Fuel * float32(tic/time.Second) / 3600.0

			if len(player.Queues) > 0 {
				if player.Queues[0].TimeLeft < 0 {
					player.MakeBuilding(player.Queues[0].Coords, player.Queues[0].Construction)
					player.Queues = append(player.Queues[:0], player.Queues[1:]...)
				} else {
					player.Queues[0].TimeLeft -= int(tic / time.Second)
					player.Queues[0].Completed = 100.0 - float32(player.Queues[0].TimeLeft*100)/float32(BUILDING_INFO[player.Queues[0].Construction.Type].BuildTime[player.Queues[0].Construction.Level])
				}
			}

			var clock Clock
			clock.Hour, clock.Min, clock.Second = time.Now().Clock()
			if clock == MidNight {
				if len(player.Debt) > 0 {
					for index, _ := range player.Debt {
						if player.Debt[index].PaymentLeft < player.Debt[index].EDI {
							player.DeductResources(Resources{player.Debt[index].PaymentLeft, 0, 0, 0, 0, 0, 0})
							player.Debt = append(player.Debt[:index], player.Debt[index+1:]...)
						} else {
							player.DeductResources(Resources{player.Debt[index].EDI, 0, 0, 0, 0, 0, 0})
							player.Debt[index].PaymentLeft -= player.Debt[index].EDI
							player.Debt[index].PaymentLeft += player.Debt[index].PaymentLeft * (player.Debt[index].Rate / 100.0)
						}
						//player.Debt[index]
					}
				}
			}

			player.Update()
		}
	}
}

func PushBuildQueue(player Player, Coords Coord, Build Building) OutMessage {
	var Message OutMessage
	var queue Queue

	if len(player.Queues) > 4 {
		Message.Check = "false"
		Message.Message = "Queue Limit Reached. Wait for initial task to complete"
	} else {
		if player.isResourceFull(BUILDING_INFO[Build.Type].Cost[Build.Level]) {
			player.DeductResources(BUILDING_INFO[Build.Type].Cost[Build.Level])
			Message.Check = "true"
			Message.Message = "Building successfully added to the queue"
			queue.Coords = Coords
			queue.Construction = Build
			queue.Completed = 0
			queue.TimeLeft = BUILDING_INFO[Build.Type].BuildTime[Build.Level]
			player.Queues = append(player.Queues, queue)
		} else {
			Message.Check = "false"
			Message.Message = "Not enough resources"
		}
	}
	player.Update()
	return Message
}
