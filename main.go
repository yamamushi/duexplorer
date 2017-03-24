
package main

import (
	"./config"
	"./models"
	"fmt"
	"log"
	"strconv"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
)



func main(){

	config.Read()
	//fmt.Println(config.GetMongoCreds())

	// There is another (probably better) way of assembling this session, but this will do for now...
	session, err := mgo.Dial("mongodb://"+config.GetMongoUser()+":"+config.GetMongoPass()+
		"@"+config.GetMongoHost()+":"+strconv.Itoa(config.GetMongoPort())+"/"+config.GetDBName())
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(config.GetDBName()).C("users")

	result := []models.Users{}
	err = c.Find(bson.M{"user":"Wardion2000"}).All(&result)
	if err != nil {
		log.Fatal(err)
	}

	formatter := json.MarshalIndent
	response, err := formatter(result, " ", "    ")

	//m, _ := json.Marshal(result)
	//t := json.Unmarshal(m, &[]models.Users{})
	fmt.Println("Pledge Status:", string(response))
	
}

