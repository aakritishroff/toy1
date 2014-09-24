package main

import "github.com/aakritishroff/toy1/toy1"

/* 
Create new mongoDB session, using database "toy1". 
Create new server for that session, and begin listening for HTTP requests.
*/

func main(){
	session := toy1.NewSession("resources")
	server := toy1.NewServer(session)
	server.Run()
}