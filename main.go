package main

import (
    "github.com/julienschmidt/httprouter"
    "gopkg.in/mgo.v2"
    "net/http"
    "restapi/controllers"
)

func main() {
  r := httprouter.New()
  uc := controllers.NewUserController(getSession())
  r.GET("/user/:id", uc.GetUser)
  r.POST("/user", uc.CreateUser)
  r.DELETE("/user/:id", uc.DeleteUser)
    err := http.ListenAndServe("localhost:9000", r)
    if err != nil {
        return 
    }
}

func getSession() *mgo.Session {
    s, err := mgo.Dial("mongodb://127.0.0.1:27017")
    if err != nil {
        panic(err)
    }
    return s
}
