package main

import (
	"github.com/yamamushi/duexplorer/config"
	"github.com/yamamushi/duexplorer/models"
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"strconv"
)

func main() {

	config.Read()
	//fmt.Println(config.GetMongoCreds())

	// There is another (probably better) way of assembling this session, but this will do for now...
	session, err := mgo.Dial("mongodb://" + config.GetMongoUser() + ":" + config.GetMongoPass() +
		"@" + config.GetMongoHost() + ":" + strconv.Itoa(config.GetMongoPort()) + "/" + config.GetDBName())
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(config.GetDBName()).C("users")
	user_result := []models.User{}
	err = c.Find(bson.M{}).All(&user_result)
	if err != nil {
		log.Fatal(err)
	}

	// Format our output properly
	formatter := json.MarshalIndent
	user_response, err := formatter(user_result, " ", "   ")

	//m, _ := json.Marshal(result)
	//t := json.Unmarshal(m, &[]models.Users{})
	fmt.Println("Users:", string(user_response))


	d := session.DB(config.GetDBName()).C("orgs")
	group_result := []models.Group{}
	err = d.Find(bson.M{}).All(&group_result)
	if err != nil {
		log.Fatal(err)
	}
	group_formatter := json.MarshalIndent
	group_response, err := group_formatter(group_result, " ", "   ")
	fmt.Println("Groups: ", string(group_response))
}
