package router

import (
	"github.com/alseRokachev/chat-app/internal/user"
	"github.com/gin-gonic/gin"
	"log"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler) {
	r = gin.Default()

	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)
}

func Start(addr string) error {
	log.Println(addr)
	return r.Run(addr)
}
