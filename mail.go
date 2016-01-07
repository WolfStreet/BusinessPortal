package main

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Mail struct {
	ID       bson.ObjectId `bson:"_id" json:"ID"`
	Sender   string        `bson:"Sender" json:"Sender"`
	Receiver []string      `bson:"Receiver" json:"Receiver"`
	Subject  string        `bson:"Subject" json:"Subject"`
	Content  string        `bson:"Content" json:"Content"`
	InOrOut  string        `bson:"InOrOut" json:"InOrOut"`
	Time     time.Time     `bson:"Time" json:"Time"`
}

func (mail *Mail) Init() {
	mail.ID = bson.NewObjectIdWithTime(time.Now())
	mail.Sender = ""
	mail.Receiver = nil
	mail.Subject = ""
	mail.Content = ""
	mail.InOrOut = ""
	mail.Time = time.Now()
}

func (mail *Mail) InitWithValues(sender string, reciever []string, subject, content, inoroout string, timeinit time.Time) {
	mail.ID = bson.NewObjectIdWithTime(time.Now())
	mail.Sender = sender
	mail.Receiver = reciever
	mail.Subject = subject
	mail.Content = content
	mail.InOrOut = inoroout
	mail.Time = timeinit
}

func (mail *Mail) InitWithMail(mailinit Mail) {
	mail.ID = mailinit.ID
	mail.Sender = mailinit.Sender
	mail.Receiver = mailinit.Receiver
	mail.Subject = mailinit.Subject
	mail.Content = mailinit.Content
	mail.InOrOut = mailinit.InOrOut
	mail.Time = mail.Time
}
