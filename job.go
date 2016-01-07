package main

import (
	//"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Job struct {
	ID          bson.ObjectId `bson:"_id" json:"ID"`
	UserName    string        `bson:"UserName" json:"UserName"`
	Name        string        `bson:"Name" json:"Name"`
	Actioner    string        `bson:"Actioner" json:"Actioner"` //Candidate - User wants to get recruited or Employer - User wants to hire
	Career      string        `bson:"Career" json:"Career"`     //Engineer or Transport
	SkillPoints int           `bson:"SkillPoints" json:"SkillPoints"`
	Fee         int           `bson:"Fee" json:"Fee"`
	Hours       int           `bson:"Hours" json:"Hours"`
}

func FetchFullJob() []Job {
	var job []Job
	session, err := mgo.DialWithInfo(MongoURL)
	CheckErr("Session error of FetchFullJob", err)
	err = session.DB("RTS").C("Jobs").Find(bson.M{}).All(&job)
	CheckErr("Find error of FetchFullJob", err)
	return job
}

func FetchJobByAction(Actioner string) []Job {
	var job []Job
	session, err := mgo.DialWithInfo(MongoURL)
	CheckErr("Session error of FetchJobByAction", err)
	err = session.DB("RTS").C("Jobs").Find(bson.M{"Actioner": Actioner}).All(&job)
	CheckErr("Document error of FetchJobByAction", err)
	return job
}

func FetchJobByCareer(Career string) []Job {
	var job []Job
	session, err := mgo.DialWithInfo(MongoURL)
	CheckErr("Session error of FetchJobByCareer", err)
	err = session.DB("RTS").C("Jobs").Find(bson.M{"Career": Career}).All(&job)
	CheckErr("Document error of FetchJobByCareer", err)
	return job
}

func (job *Job) FetchJobByID() {
	session, err := mgo.DialWithInfo(MongoURL)
	CheckErr("Session error of FetchJobByID", err)
	err = session.DB("RTS").C("Jobs").Find(bson.M{"_id": job.ID}).One(&job)
	CheckErr("Document error of FetchJobByID", err)
}

func (job *Job) SaveJob() {
	session, err := mgo.DialWithInfo(MongoURL)
	defer session.Close()
	CheckErr("Session error of SaveJob", err)
	err = session.DB("RTS").C("Jobs").Update(bson.M{"_id": job.ID}, &job)
	CheckErr("Update error of SaveJob", err)
}

func (job *Job) PostJob() OutMessage {
	var player Player
	var Message OutMessage
	session, err := mgo.DialWithInfo(MongoURL)
	defer session.Close()
	CheckErr("Session error of PostJob", err)
	err = session.DB("RTS").C("PlayerData").Find(bson.M{"UserName": job.UserName}).One(&player)
	CheckErr("Player error of PostJob", err)
	job.Name = player.Name
	job.SkillPoints = player.SkillPoints
	job.Career = player.Career
	if job.Actioner == "Candidate" {
		if !(job.Career == CAREER_INDEX_ENGINEER || job.Career == CAREER_INDEX_TRANSPORT) {
			Message.Check = "false"
			Message.Message = "Only Engineer and Transport can post job"
			return Message
		}
	}
	Message.Check = "true"
	Message.Message = "Job Request Posted"

	err = session.DB("RTS").C("Jobs").Insert(&job)
	CheckErr("Insert error of PostJob", err)
	return Message
}

func (job *Job) AcceptJob(player Player) OutMessage {
	var Message OutMessage
	return Message
}
