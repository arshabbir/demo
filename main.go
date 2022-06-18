package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type server struct {
}

type Server interface {
	HandlePing(c *gin.Context)
	HandleUsers(c *gin.Context)
	HandleUserActions(c *gin.Context)
}

func main() {

	r := gin.Default()

	addr := ":8080"
	s := NewServer()

	r.GET("/ping", s.HandlePing)

	// /users/:id
	r.POST("/users/:id", s.HandleUsers)

	// /users/:id/action
	r.POST("/users", s.HandleUsers)

	r.GET("/favicon.ico", nil)

	if err := r.Run(addr); err != nil {
		log.Println("error while running", err)
	}

}

func NewServer() Server {
	return &server{}
}
func (s *server) HandlePing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (s *server) HandleUsers(c *gin.Context) {
	id := c.Param("id")

	// firstname := c.Query("firstname")
	// lastname := c.DefaultQuery("lastname", "Guest")
	firstname := c.PostForm("firstname")
	lastname := c.PostForm("lastname")

	// if id == "" {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "id cannot be nil",
	// 	})
	// 	return
	// }
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Response : %s FirstName : %s, Lastname :  %s", id, firstname, lastname),
	})

}

func (s *server) HandleUserActions(c *gin.Context) {
	id := c.Param("id")
	action := c.Param("action")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id cannot be nil",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Response : %s - %s", id, action),
	})

}