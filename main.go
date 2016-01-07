//package BusinessPortal
package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

type OutMessage struct {
	Check   string `bson:"Check" json:"Check"`
	Message string `bson:"Message" json:"Message"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/Login.html")
	CheckErr("Parsing error of LoginHandler", err)
	t.Execute(w, nil)
}

func CheckLoginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	var Message OutMessage
	user.UserName = r.PostFormValue("UserName")
	temp := r.PostFormValue("Password")
	user.Password = EncryptData(temp)
	ValidLogin := ValidUser(user)
	if ValidLogin == true {
		Message.Check = "true"
		Message.Message = "Valid Login"
	} else {
		Message.Check = "false"
		Message.Message = "Invalid Login! Please Enter Valid Details."
	}

	MessageJSON, _ := json.Marshal(&Message)
	w.Write(MessageJSON)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/Register.html")
	CheckErr("Parsing error of RegisterHandler", err)
	t.Execute(w, nil)

}

func CheckUserHandler(w http.ResponseWriter, r *http.Request) {
	UserName := r.PostFormValue("UserName")
	UserExist := FetchUser(UserName)
	var Message OutMessage

	if UserExist == true {
		Message.Check = "false"
		Message.Message = "User Already Exist Please Try Another Name."
	} else {
		Message.Check = "true"
		Message.Message = "User Name Available. Go Ahead and Register!"
	}

	MessageJSON, _ := json.Marshal(&Message)
	w.Write(MessageJSON)
}

func SaveUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	user.Id = bson.NewObjectId()
	user.UserName = r.PostFormValue("UserName")
	user.Name = r.PostFormValue("Name")
	temp := r.PostFormValue("Password")
	user.Password = EncryptData(temp)
	SaveUser(user)
}

func WorldHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		UserName := r.PostFormValue("UserName")
		player := FetchPlayerByName(UserName)
		t, err := template.ParseFiles("./templates/World.html")
		if err != nil {
			fmt.Println(err)
		}
		t.Execute(w, player)
	} else {
		http.Error(w, "Unable to fetch the player. Please login again.", 200)
	}
}

func FetchPlayerHandler(w http.ResponseWriter, r *http.Request) {
	var UserName = r.PostFormValue("UserName")
	player := FetchPlayerByName(UserName)
	playerJSON, _ := json.Marshal(&player)
	w.Write([]byte(playerJSON))
}

func SavePlayerHandler(w http.ResponseWriter, r *http.Request) {
	var player Player
	r.ParseForm()
	var FirstPlay string
	FirstPlay = r.PostFormValue("FirstPlay")
	Career := r.PostFormValue("Career")
	UserName := r.PostFormValue("UserName")
	player = FetchPlayerByName(UserName)
	if FirstPlay == "FirstPlay" {
		player.InitPlayer(UserName, player.Name, Career)
		player.Update()
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&player)
	if err != nil {
		fmt.Println(err)
	}
	player.Update()
}

func WorldBuildingHandler(w http.ResponseWriter, r *http.Request) {
	var Build Building
	var Message OutMessage
	var UserName = r.PostFormValue("UserName")
	var Coords Coord
	Coords.X, _ = strconv.Atoi(r.FormValue("CoordsX"))
	Coords.Y, _ = strconv.Atoi(r.FormValue("CoordsY"))
	Build.Type, _ = strconv.Atoi(r.FormValue("BuildType"))
	Build.Level, _ = strconv.Atoi(r.FormValue("BuildLevel"))

	Message = PushBuildQueue(FetchPlayerByName(UserName), Coords, Build)

	MessageJSON, _ := json.Marshal(&Message)
	w.Write(MessageJSON)
}

func CancelQueueHandler(w http.ResponseWriter, r *http.Request) {
	var Message OutMessage
	var Coords Coord
	var Build Building

	var UserName = r.PostFormValue("UserName")
	player := FetchPlayerByName(UserName)
	Coords.X, _ = strconv.Atoi(r.FormValue("CoordsX"))
	Coords.Y, _ = strconv.Atoi(r.FormValue("CoordsY"))
	Build.Type, _ = strconv.Atoi(r.FormValue("BuildType"))
	Build.Level, _ = strconv.Atoi(r.FormValue("BuildLevel"))

	Message = player.CancelQueue(Coords, Build)
	MessageJSON, _ := json.Marshal(&Message)
	w.Write(MessageJSON)
}

func BuildingInfoHandler(w http.ResponseWriter, r *http.Request) {
	return
}

func CareerRequestHandler(w http.ResponseWriter, r *http.Request) {
	var player []Player
	career := r.PostFormValue("Career")
	player = FetchPlayerByCareer(career)

	playerJSON, _ := json.Marshal(&player)
	w.Write(playerJSON)
}

/*---Post Handler Begins---*/
func PostBankHandler(w http.ResponseWriter, r *http.Request) {
	var bank Bank
	var Message OutMessage
	bank.ID = bson.NewObjectIdWithTime(time.Now())
	bank.UserName = r.PostFormValue("UserName")
	bank.Actioner = r.PostFormValue("Actioner")
	bank.Amount, _ = strconv.Atoi(r.PostFormValue("Amount"))
	bank.Rate, _ = strconv.Atoi(r.PostFormValue("Rate"))
	Message = bank.PostBank()
	MessageJSON, _ := json.Marshal(&Message)
	w.Write(MessageJSON)
}

func PostJobsHandler(w http.ResponseWriter, r *http.Request) {
	var job Job
	var Message OutMessage
	job.ID = bson.NewObjectIdWithTime(time.Now())
	job.UserName = r.PostFormValue("UserName")
	job.Actioner = r.PostFormValue("Actioner")
	job.Fee, _ = strconv.Atoi(r.PostFormValue("Fee"))
	job.Hours, _ = strconv.Atoi(r.PostFormValue("Hours"))
	Message = job.PostJob()
	MessageJSON, _ := json.Marshal(&Message)
	w.Write(MessageJSON)
}

func PostMarketHandler(w http.ResponseWriter, r *http.Request) {
	var market Market
	var Message OutMessage
	market.ID = bson.NewObjectIdWithTime(time.Now())
	market.UserName = r.PostFormValue("UserName")
	market.Name = r.PostFormValue("Name")
	market.Actioner = r.PostFormValue("Actioner")
	market.Resource = r.PostFormValue("Resource")
	market.Rate, _ = strconv.Atoi(r.PostFormValue("Rate"))
	market.Amount, _ = strconv.Atoi(r.PostFormValue("Amount"))
	Message = market.PostMarket()
	MessageJSON, _ := json.Marshal(&Message)
	w.Write(MessageJSON)
}

func PostMailHandler(w http.ResponseWriter, r *http.Request) {

}

/*---Post Handler Ends---*/
/*---Fetch Handler Begins---*/
func FetchBankHandler(w http.ResponseWriter, r *http.Request) {
	var bank []Bank
	var temp Bank
	temp.UserName = r.PostFormValue("Users")
	temp.Actioner = r.PostFormValue("Actioner")
	if temp.UserName == "All" || temp.UserName == "all" {
		bank = FetchBankbyAction(temp.Actioner)
		bankJSON, _ := json.Marshal(&bank)
		w.Write(bankJSON)
	} else {
		temp.ID = bson.ObjectIdHex(r.PostFormValue("ID"))
		temp.FetchBankByID()
		bankJSON, _ := json.Marshal(&temp)
		w.Write(bankJSON)
	}
}

func FetchJobsHandler(w http.ResponseWriter, r *http.Request) {
	var job []Job
	var temp Job
	temp.UserName = r.PostFormValue("Users")
	temp.Actioner = r.PostFormValue("Actioner")
	temp.Career = r.PostFormValue("Career")
	if temp.UserName == "All" || temp.UserName == "all" {
		if temp.Career != "" {
			job = FetchJobByCareer(temp.Career)
		} else {
			job = FetchJobByAction(temp.Actioner)
		}
		jobsJSON, _ := json.Marshal(&job)
		w.Write(jobsJSON)
	} else {
		temp.ID = bson.ObjectIdHex(r.PostFormValue("ID"))
		temp.FetchJobByID()
		jobsJSON, _ := json.Marshal(&temp)
		w.Write(jobsJSON)
	}
}

func FetchMarketHandler(w http.ResponseWriter, r *http.Request) {
	var market []Market
	var temp Market
	temp.UserName = r.PostFormValue("Users")
	temp.Actioner = r.PostFormValue("Actioner")
	if temp.UserName == "All" || temp.UserName == "all" {
		market = FetchMarketByAction(temp.Actioner)
		marketJSON, _ := json.Marshal(&market)
		w.Write(marketJSON)
	} else {
		temp.ID = bson.ObjectIdHex(r.PostFormValue("ID"))
		temp.FetchMarketByID()
		marketJSON, _ := json.Marshal(&temp)
		w.Write(marketJSON)
	}
}

func FetchMailHandler(w http.ResponseWriter, r *http.Request) {

}

/*---Fetch Handler Ends---*/
/*---Accept Handler Begins---*/
func AcceptBankHandler(w http.ResponseWriter, r *http.Request) {
	var bank Bank
	var player Player
	var Message OutMessage
	if bson.IsObjectIdHex(r.PostFormValue("ID")) {
		bank.ID = bson.ObjectIdHex(r.PostFormValue("ID"))
	} else {
		fmt.Println(r.PostFormValue("ID"))
		Message.Check = "false"
		Message.Message = "Can't find the Bank ID. Please report the bug."
		MessageJSON, _ := json.Marshal(&Message)
		w.Write(MessageJSON)
		return
	}
	player.UserName = r.PostFormValue("UserName")
	player = FetchPlayerByName(player.UserName)
	Message = bank.AcceptBank(player)
	MessageJSON, _ := json.Marshal(&Message)
	w.Write(MessageJSON)
}

func AcceptJobHandler(w http.ResponseWriter, r *http.Request) {
	var job Job
	var player Player
	var Message OutMessage
	if bson.IsObjectIdHex(r.PostFormValue("ID")) {
		job.ID = bson.ObjectIdHex(r.PostFormValue("ID"))
	} else {
		Message.Check = "false"
		Message.Message = "Can't find the Job ID. Please report the bug."
		MessageJSON, _ := json.Marshal(&Message)
		w.Write(MessageJSON)
		return
	}
	player.UserName = r.PostFormValue("UserName")
	player = FetchPlayerByName(player.UserName)
	Message = job.AcceptJob(player)
}

func AcceptMarketHandler(w http.ResponseWriter, r *http.Request) {
	var market Market
	var player Player
	var Message OutMessage
	if bson.IsObjectIdHex(r.PostFormValue("ID")) {
		market.ID = bson.ObjectIdHex(r.PostFormValue("ID"))
	} else {
		Message.Check = "false"
		Message.Message = "Can't find the Market ID. Please report the bug."
		MessageJSON, _ := json.Marshal(&Message)
		w.Write(MessageJSON)
		return
	}
	player.UserName = r.PostFormValue("UserName")
	player = FetchPlayerByName(player.UserName)
	Message = market.AcceptMarket(player)
	MessageJSON, _ := json.Marshal(&Message)
	w.Write(MessageJSON)
}

func AcceptMailHandler(w http.ResponseWriter, r *http.Request) {

}

/*---Accept Handler Ends---*/
func main() {

	go CatchTic()
	mgo.SetLogger(Logger)

	var r = mux.NewRouter()
	r.StrictSlash(true)
	r.HandleFunc("/", LoginHandler)
	r.HandleFunc("/login/", CheckLoginHandler)
	r.HandleFunc("/register/", RegisterHandler)
	r.HandleFunc("/register/check/", CheckUserHandler)
	r.HandleFunc("/register/save/", SaveUserHandler)

	r.HandleFunc("/world01/", WorldHandler)
	r.HandleFunc("/world01/fetch/player/", FetchPlayerHandler)
	r.HandleFunc("/world01/save/player/", SavePlayerHandler)

	r.HandleFunc("/world01/building/build/", WorldBuildingHandler)
	r.HandleFunc("/world01/building/cancelQueue/", CancelQueueHandler)

	r.HandleFunc("/world01/post/bank/", PostBankHandler)
	r.HandleFunc("/world01/post/job/", PostJobsHandler)
	r.HandleFunc("/world01/post/market/", PostMarketHandler)
	r.HandleFunc("/world01/post/mail/", PostMailHandler)

	r.HandleFunc("/world01/fetch/bank/", FetchBankHandler)
	r.HandleFunc("/world01/fetch/job/", FetchJobsHandler)
	r.HandleFunc("/world01/fetch/market/", FetchMarketHandler)
	r.HandleFunc("/world01/fetch/mail/", FetchMailHandler)

	r.HandleFunc("/world01/accept/bank/", AcceptBankHandler)
	r.HandleFunc("/world01/accept/job/", AcceptJobHandler)
	r.HandleFunc("/world01/accept/market/", AcceptMarketHandler)
	r.HandleFunc("/world01/accept/market/", AcceptMailHandler)

	r.HandleFunc("/players/career/", CareerRequestHandler)

	r.HandleFunc("/building/info/", BuildingInfoHandler)

	r.PathPrefix("/style/").Handler(http.StripPrefix("/style/", http.FileServer(http.Dir("./style/"))))
	r.PathPrefix("/script/").Handler(http.StripPrefix("/script/", http.FileServer(http.Dir("./script/"))))
	r.PathPrefix("/src/").Handler(http.StripPrefix("/src/", http.FileServer(http.Dir("./src/"))))
	r.PathPrefix("/templates/").Handler(http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates/"))))
	http.Handle("/", r)

	IPAddress := "192.168.0.5" //"127.0.0.1" //10.101.7.85
	Port := ":90"
	go http.ListenAndServe(":80", http.RedirectHandler("https://192.168.0.5:90/", 301))
	err := http.ListenAndServeTLS(IPAddress+Port, "cert.pem", "key.pem", nil)
	CheckErr("ListenandServe error", err)
}
