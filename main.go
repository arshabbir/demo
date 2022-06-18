package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type server struct {
}

type Server interface {
	HandlePing(c *gin.Context)
	HandleUsers(c *gin.Context)
	HandleUserActions(c *gin.Context)
	HandleQueryMap(c *gin.Context)
	HandleLogin(c *gin.Context)
	HandleLoginV2(c *gin.Context)
	LoggerMiddleware() gin.HandlerFunc
	AuthMiddleware() gin.HandlerFunc
}

func main() {

	r := gin.New()

	addr := ":8080"
	s := NewServer()

	r.GET("/ping", s.HandlePing)

	// /users/:id
	r.POST("/users/:id", s.HandleUsers)

	r.POST("/users", s.HandleUsers)

	r.GET("/favicon.ico", nil)

	r.GET("/querymap", s.HandleQueryMap)

	v1 := r.Group("/v1")
	// Middleware
	v1.Use(s.LoggerMiddleware())
	v1.Use(s.AuthMiddleware())

	v2 := r.Group("/v2")

	v1.GET("/login", s.HandleLogin)
	v2.GET("/login", s.HandleLoginV2)

	if err := r.Run(addr); err != nil {
		log.Println("error while running", err)
	}

}

func NewServer() Server {
	return &server{}
}

func (s *server) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Auth middleware invoked")
		c.Header("id", fmt.Sprintf("%v", time.Now().UnixNano()))
		c.Next()
	}
}

func (s *server) LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("logging middleware invoked.")
		c.Next()
	}
}
func (s *server) HandleLogin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "v1 login",
	})

}

func (s *server) HandleLoginV2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "v2 login",
	})

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

func (s *server) HandleQueryMap(c *gin.Context) {

	m := c.QueryMap("users")
	log.Println(m)
	c.JSON(http.StatusOK, m)

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
