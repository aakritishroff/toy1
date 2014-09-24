package toy1

import "github.com/go-martini/martini"
import "gopkg.in/mgo.v2"

/*
By starting a DBSession, I can use different database tables for tests, and additions later on. 
*/
type DatabaseSession struct {
	*mgo.Session 
	dbName string
}

/*
Connect to local mongoDB and setup DB
*/
func NewSession(dbName string) *DatabaseSession {
	session, err := mgo.Dial("mongodb://localhost")
    if err != nil {
        panic(err)
    }

    return &DatabaseSession{session, dbName}
}

/*
context.Map() gets an instance of the *mgo.Database to get and post resources
*/
func (session *DatabaseSession) Database() martini.Handler {
	return func(context martini.Context){
		s := session.Clone()
		context.Map(s.DB(session.dbName))
		defer s.Close()
		context.Next()
	}
}
