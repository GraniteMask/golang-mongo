package main

import (
"github.com/julienschmidt/httprouter"
"gopkg.in/mgo.v2"
"net/http"
"appointy-Golang-mongo/controllers"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/users/:id", uc.GetUser)
	r.POST("/users", uc.CreateUser)
	r.GET("/posts/:id", uc.GetPost)
	r.POST("/posts", uc.CreatePost)
	http.ListenAndServe("localhost:9000", r)
}

func getSession() *mgo.Session{
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err !=nil{
		panic(err)
	}
	return s
}

// mongodb+srv://ratnadeep:CpXnwKKpQNYnNrwG@cluster0.d1fon.mongodb.net/myFirstDatabase?retryWrites=true&w=majority

//mongodb://localhost:27107