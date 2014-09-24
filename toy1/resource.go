package toy1

import "gopkg.in/mgo.v2"
import "fmt"
//import "encoding/json"
//import uuid "github.com/satori/go.uuid"

type Resource struct {
	Id  int `json:"Id"`
	Item interface{} `json:"Item"`
}


func fetchAll(db *mgo.Database) []Resource {
	resources := []Resource{}
	err := db.C("resources").Find(nil).All(&resources)
	if err != nil {
		fmt.Println("errored out")
		panic(err)
	}
	return resources
}