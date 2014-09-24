package toy1

import (
    "github.com/go-martini/martini"
    "github.com/martini-contrib/binding"
    "github.com/martini-contrib/render"
   	"gopkg.in/mgo.v2"
   	"fmt"
)

/*
Martini server struct wrapper
*/ 
type Server *martini.ClassicMartini

func NewServer(session *DatabaseSession) Server {
	m := Server(martini.Classic())
	m.Use(render.Renderer(render.Options{
		IndentJSON: true,
		}))
	m.Use(session.Database())

	// Define "GET /public" route
	m.Get("/toy1/public", func(r render.Render, db *mgo.Database) {
		//session.DB("resources").DropDatabase()
		r.JSON(200, fetchAll(db))
	})

	// Define "POST /public" route
	m.Post("/toy1/public", binding.Json(Request{}), 
		func(request Request, 
			r render.Render, 
			db *mgo.Database) {
			fmt.Printf("%+v", request.Items)
			if request.validReq(){
				err := request.parseReq(db)
				if err == nil {
					r.JSON(201, request)
				} else {
					r.JSON(400, map[string]string{
						"error": err.Error(),
					})
				}
			} else {
				r.JSON(400, map[string]string{
					"error":"Not a valid resource",
				})
			}
		})

	return m
}