//package BusinessPortal
package main

import (
	//"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Friend string

type Coord struct {
	X int `bson:"X" json:"X"`
	Y int `bson:"Y" json:"Y"`
}

type Employee struct {
	UserName    string `bson:"UserName" json:"UserName"`
	Career      string `bson:"Career" json:"Career"`
	Fee         int    `bson:"Fee" json:"Fee"`
	Duration    int    `bson:"Duration" json:"Duration"`
	SkillPoints int    `bson:"SkillPoints" json:"SkillPoints"`
}

type Queue struct {
	Construction Building `bson:"Construction" json:"Construction"`
	Completed    float32  `bson:"Completed" json:"Completed"`
	TimeLeft     int      `bson:"TimeLeft" json:"TimeLeft"`
	Coords       Coord    `bson:"Coords" json:"Coords"`
}

type Loan struct {
	Amount      float32 `bson:"Amount" json:"Amount"`
	Rate        float32 `bson:"Rate" json:"Rate"`
	Lender      string  `bson:"Lender" json:"Lender"`
	EDI         float32 `bson:"EDI" json:"EDI"`
	PaymentLeft float32 `bson:"PaymentLeft" json:"PaymentLeft"`
}

type Player struct {
	UserName    string `bson:"UserName" json:"UserName"`
	Name        string `bson:"Name" json:"Name"`
	Career      string `bson:"Career" json:"Career"`
	SkillPoints int    `bson:"SkillPoints" json:"SkillPoints"`
	Friends     []Friend

	Mails []Mail

	ResStored     Resources `bson:"ResStored" json:"ResStored"`
	ResProduction Resources `bson:"ResProduction" json:"ResProduction"`

	Buildings [10][10]Building `bson:"Buildings" json:"Buildings"`
	Queues    []Queue          `bson:"Queues" json:"Queues"`

	Debt      []Loan     `bson:"Debt" json:"Debt"`
	Employees []Employee `bson:"Employees" json:"Employees"`
}

const (
	CAREER_INDEX_BANKER       = "Banker"
	CAREER_INDEX_DEALER       = "Dealer"
	CAREER_INDEX_TRANSPORT    = "Transport"
	CAREER_INDEX_ENGINEER     = "Engineer"
	CAREER_INDEX_ENTREPRENEUR = "Entrepreneur"
)

func (player *Player) InitPlayer(UserName string, Name string, Career string) {

	player.UserName = UserName
	player.Name = Name
	player.Career = Career

	switch Career {
	case CAREER_INDEX_BANKER:
		player.ResStored = DEFAULT_RESOURCES_BANKER
		player.ResProduction = DEFAULT_PRODUCTION_BANKER
		break
	case CAREER_INDEX_DEALER:
		player.ResStored = DEFAULT_RESOURCES_DEALER
		player.ResProduction = DEFAULT_PRODUCTION_DEALER
		break
	case CAREER_INDEX_TRANSPORT:
		player.ResStored = DEFAULT_RESOURCES_TRANSPORT
		player.ResProduction = DEFAULT_PRODUCTION_TRANSPORT
		break
	case CAREER_INDEX_ENGINEER:
		player.ResStored = DEFAULT_RESOURCES_ENGINEER
		player.ResProduction = DEFAULT_PRODUCTION_ENGINEER
		break
	case CAREER_INDEX_ENTREPRENEUR:
		player.ResStored = DEFAULT_RESOURCES_ENTREPRENEUR
		player.ResProduction = DEFAULT_PRODUCTION_ENTREPRENEUR
		break
	}

	for index_i, building := range player.Buildings {
		for index_j, _ := range building {
			player.Buildings[index_i][index_j].Type = 0
			player.Buildings[index_i][index_j].Level = 0
		}
	}

	player.SkillPoints = 5

	var queue Queue
	queue.Completed = 100
	queue.TimeLeft = 0
	queue.Construction.Type = 0
	queue.Construction.Level = 0
	queue.Coords.X = 0
	queue.Coords.Y = 0
	player.Queues = append(player.Queues, queue)

	var debt Loan
	debt.Amount = 0.0
	debt.Rate = 0.0
	debt.EDI = 0
	debt.PaymentLeft = 0
	debt.Lender = ""
	player.Debt = append(player.Debt, debt)

	var friend Friend
	friend = ""
	player.Friends = append(player.Friends, friend)

	var mail Mail
	mail.Init()
	player.Mails = append(player.Mails, mail)
}

func (player *Player) Save() {
	session, err := mgo.DialWithInfo(MongoURL)
	defer session.Close()
	CheckErr("Session error of Player Save", err)
	err = session.DB("RTS").C("PlayerData").Insert(&player)
	CheckErr("Insert error of Player Save", err)
}

func (player *Player) Update() {
	session, err := mgo.DialWithInfo(MongoURL)
	defer session.Close()
	CheckErr("Session error of Player Update", err)
	err = session.DB("RTS").C("PlayerData").Update(bson.M{"UserName": player.UserName}, &player)
	CheckErr("Update error of Player Update", err)
}

func (player *Player) MakeBuilding(Coords Coord, Build Building) {
	player.Buildings[Coords.X][Coords.Y].Type = Build.Type
	player.Buildings[Coords.X][Coords.Y].Level = Build.Level
	player.UpdateProduction()
	player.Update()
}

func (player *Player) UpdateProduction() {
	var career int
	switch player.Career {
	case CAREER_INDEX_BANKER:
		career = 0
		player.ResProduction = DEFAULT_PRODUCTION_BANKER
		break
	case CAREER_INDEX_DEALER:
		career = 1
		player.ResProduction = DEFAULT_PRODUCTION_DEALER
		break
	case CAREER_INDEX_TRANSPORT:
		career = 2
		player.ResProduction = DEFAULT_PRODUCTION_TRANSPORT
		break
	case CAREER_INDEX_ENGINEER:
		career = 3
		player.ResProduction = DEFAULT_PRODUCTION_ENGINEER
		break
	case CAREER_INDEX_ENTREPRENEUR:
		player.SkillPoints = 0
		career = 4
		player.ResProduction = DEFAULT_PRODUCTION_ENTREPRENEUR
		break
	}

	for _, row := range player.Buildings {
		for _, col := range row {
			switch col.Type {
			case ASSET_INDEX_HOUSE:
				player.ResProduction.Money += BUILDING_INFO[ASSET_INDEX_HOUSE].Production[col.Level] * BUILDING_PRODUCTION_FACTOR[career].Money
				if player.isEntrepreneur() {
					player.SkillPoints += 1 * col.Level
				}
				break
			case ASSET_INDEX_FLAT:
				player.ResProduction.Money += BUILDING_INFO[ASSET_INDEX_FLAT].Production[col.Level] * BUILDING_PRODUCTION_FACTOR[career].Money
				if player.isEntrepreneur() {
					player.SkillPoints += 1 * col.Level
				}
				break
			case ASSET_INDEX_SHOP:
				player.ResProduction.Money += BUILDING_INFO[ASSET_INDEX_SHOP].Production[col.Level] * BUILDING_PRODUCTION_FACTOR[career].Money
				if player.isEntrepreneur() {
					player.SkillPoints += 1 * col.Level
				}
				break
			case GENERATOR_INDEX_FOOD:
				player.ResProduction.Food += BUILDING_INFO[GENERATOR_INDEX_FOOD].Production[col.Level] * BUILDING_PRODUCTION_FACTOR[career].Food
				if player.isEntrepreneur() {
					player.SkillPoints += 1 * col.Level
				}
				break
			case GENERATOR_INDEX_WOOD:
				player.ResProduction.Wood += BUILDING_INFO[GENERATOR_INDEX_WOOD].Production[col.Level] * BUILDING_PRODUCTION_FACTOR[career].Wood
				if player.isEntrepreneur() {
					player.SkillPoints += 1 * col.Level
				}
				break
			case GENERATOR_INDEX_METAL:
				player.ResProduction.Metal += BUILDING_INFO[GENERATOR_INDEX_METAL].Production[col.Level] * BUILDING_PRODUCTION_FACTOR[career].Metal
				if player.isEntrepreneur() {
					player.SkillPoints += 1 * col.Level
				}
				break
			case GENERATOR_INDEX_CEMENT:
				player.ResProduction.Cement += BUILDING_INFO[GENERATOR_INDEX_CEMENT].Production[col.Level] * BUILDING_PRODUCTION_FACTOR[career].Cement
				if player.isEntrepreneur() {
					player.SkillPoints += 1 * col.Level
				}
				break
			case GENERATOR_INDEX_CLOTHES:
				player.ResProduction.Clothes += BUILDING_INFO[GENERATOR_INDEX_CLOTHES].Production[col.Level] * BUILDING_PRODUCTION_FACTOR[career].Clothes
				if player.isEntrepreneur() {
					player.SkillPoints += 1 * col.Level
				}
				break
			case GENERATOR_INDEX_FUEL:
				player.ResProduction.Fuel += BUILDING_INFO[GENERATOR_INDEX_FUEL].Production[col.Level] * BUILDING_PRODUCTION_FACTOR[career].Fuel
				if player.isEntrepreneur() {
					player.SkillPoints += 1 * col.Level
				}
			}
		}
	}
}
func (player *Player) isBanker() bool {
	if player.Career == "Banker" {
		return true
	}
	return false
}

func (player *Player) isDealer() bool {
	if player.Career == "Dealer" {
		return true
	}
	return false
}

func (player *Player) isTransport() bool {
	if player.Career == "Transport" {
		return true
	}
	return false
}

func (player *Player) isEngineer() bool {
	if player.Career == "Engineer" {
		return true
	}
	return false
}

func (player *Player) isEntrepreneur() bool {
	if player.Career == "Entrepreneur" {
		return true
	}
	return false
}

func (player *Player) isResourceFull(resources Resources) bool {
	if player.ResStored.Money < resources.Money {
		return false
	}

	if player.ResStored.Food < resources.Food {
		return false
	}

	if player.ResStored.Wood < resources.Wood {
		return false
	}

	if player.ResStored.Metal < resources.Metal {
		return false
	}

	if player.ResStored.Cement < resources.Cement {
		return false
	}

	if player.ResStored.Clothes < resources.Clothes {
		return false
	}

	if player.ResStored.Fuel < resources.Fuel {
		return false
	}

	return true
}

func (player *Player) DeductResources(resources Resources) {
	player.ResStored.Money -= resources.Money
	player.ResStored.Food -= resources.Food
	player.ResStored.Wood -= resources.Wood
	player.ResStored.Metal -= resources.Metal
	player.ResStored.Cement -= resources.Cement
	player.ResStored.Clothes -= resources.Clothes
	player.ResStored.Fuel -= resources.Fuel
}

func (player *Player) RefundResources(resources Resources) {
	player.ResStored.Money += resources.Money
	player.ResStored.Food += resources.Food
	player.ResStored.Wood += resources.Wood
	player.ResStored.Metal += resources.Metal
	player.ResStored.Cement += resources.Cement
	player.ResStored.Clothes += resources.Clothes
	player.ResStored.Fuel += resources.Fuel
}

func (player *Player) CancelQueue(Coords Coord, Build Building) OutMessage {
	session, err := mgo.DialWithInfo(MongoURL)
	defer session.Close()
	CheckErr("Session error of Player CancelQueue", err)
	var temp Player
	var Message OutMessage

	err = session.DB("RTS").C("PlayerData").Find(bson.M{"UserName": player.UserName}).One(&temp)
	if err != nil {
		Message.Check = "false"
		Message.Message = "No Such Player exist"
	}
	for index, _ := range player.Queues {
		if player.Queues[index].Coords == Coords && player.Queues[index].Construction == Build {
			player.Queues = append(player.Queues[:index], player.Queues[index+1:]...)
			player.RefundResources(BUILDING_INFO[Build.Type].Cost[Build.Level])
			Message.Check = "true"
			Message.Message = "Queue Removed"
			player.Update()
			return Message
		}
	}
	Message.Check = "false"
	Message.Message = "Could not find the Queue"
	player.Update()
	return Message
}

func (player *Player) PushMail(mail Mail) {
	player.Mails = append(player.Mails, mail)
	player.Update()
}
