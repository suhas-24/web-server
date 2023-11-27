package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct{
	ID string `json:"user_id"`
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `json:"email"`
}

var users = []user{
	{ID:"1", Name:"Elsa", Age:24, Email:"elsa@gmail.com"},
	{ID:"2", Name:"Lolita", Age:23, Email:"lolita@gmail.com"},
	{ID:"3", Name:"Nivi", Age:34, Email:"n1vi@gmail.com"},
	{ID:"4", Name:"Rexa", Age:19, Email:"rexa@gmail.com"},
	{ID:"5", Name:"Barlie", Age:28, Email:"barlie@gmail.com"},
}


func getUsers(c *gin.Context){
	c.IndentedJSON(http.StatusOK, users)
}

func main(){
	router := gin.Default()
	router.GET("/users",getUsers)
	router.Run("localhost:8080")
}
