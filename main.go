
package main

import (
	config "./configreader"
	"./models"
	"fmt"
	"log"
	"strconv"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/yamamushi/duexplorer/configreader"
)



func main(){

	config.Read()
	//fmt.Println(configreader.GetMongoCreds())

	// There is another (probably better) way of assembling this session, but this will do for now...
	session, err := mgo.Dial("mongodb://"+config.GetMongoUser()+":"+config.GetMongoPass()+
		"@"+config.GetMongoHost()+":"+strconv.Itoa(config.GetMongoPort())+"/"+config.GetDBName())
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	/*
	c := session.DB("test").C("people")
	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	} */


	//result := models.Users{}
	c := session.DB(configreader.GetDBName()).C("users")

	var result = models.Users{}
	err = c.Find(bson.M{"pledgeStatus": "gold"}).All(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Pledge Status:", result)
	
}