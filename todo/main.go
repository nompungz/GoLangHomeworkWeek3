package main

import "github.com/gin-gonic/gin"
import "net/http"
import "strconv"

type Todo struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Status string `json:"status"`
}

type ResponseStatus struct {
	Status string `json:"status"`
}

var default_id = 10000
var todoMap = make(map[int]Todo)

func main() {
	r := gin.Default()
	r.GET("/api/todos",getTodoListHandler)
	r.GET("/api/todos/:id",getTodosHandler)
	r.POST("/api/todos",postTodosHandler)
	r.PUT("/api/todos/:id",putTodosHandler)
	r.DELETE("/api/todos/:id",deleteTodosHandler)
	r.Run(":1234")
}

func getTodoListHandler(c *gin.Context){
	tt := []Todo{}
	for _,t := range todoMap {
		tt = append(tt,t)
	}
	c.JSON(http.StatusOK,tt)
}

func postTodosHandler(c *gin.Context){
	var json Todo
	if err := c.ShouldBindJSON(&json); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	}
	default_id++ 
	id := default_id
	json.Id = default_id
	todoMap[id] = json
	response := todoMap[id]
	c.JSON(http.StatusCreated,response)
}

func getTodosHandler(c *gin.Context){
	id,error := strconv.Atoi(c.Param("id"));
	if  error != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":error.Error()})
	}
	response := todoMap[id]
	c.JSON(http.StatusOK,response)
}

func putTodosHandler(c *gin.Context){
	id,error := strconv.Atoi(c.Param("id"))
	if error != nil {
		c.JSON(http.StatusBadRequest,error.Error())
	}
	if _,ok := todoMap[id] ; !ok {
		c.JSON(http.StatusBadRequest,"id not found")
	}

	var json Todo
	if error := c.ShouldBindJSON(&json); error != nil{
		c.JSON(http.StatusBadRequest,error.Error())
	}

	todoMap[id] = Todo{ Id : id ,Title : json.Title,  Status : json.Status}
	c.JSON(http.StatusOK,todoMap[id])
}

func deleteTodosHandler(c * gin.Context){
	id,error := strconv.Atoi(c.Param("id"))
	if error != nil {
		c.JSON(http.StatusBadRequest,error.Error())
	}

	delete(todoMap,id)
	c.JSON(http.StatusOK,ResponseStatus{Status:"success"})
}