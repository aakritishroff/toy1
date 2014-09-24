package toy1

import "gopkg.in/mgo.v2"
//import "fmt"
//import "encoding/json"
import "errors"
import "math/rand"

type Request struct {
	Op string `json:"Op"`
	Items []interface{} `json:"Items"`
}

func (request *Request) validReq() bool {
	return true
}

/*
Takes requests, parses Op, creates Resource after assigning Unique ID.
Inserts each resource into db
*/
func (request *Request) parseReq(db *mgo.Database) error {
	if request.Op == "insert" {
		for _, item := range request.Items {
			resource := Resource{Id: rand.Int(), Item: item}
			_ = db.C("resources").Insert(resource)
		}
		return nil
	} else {
		return errors.New("Incorrect Op")
	}
	
}