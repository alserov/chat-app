package router

import (
	"context"
	"github.com/alseRokachev/chat-app/internal/user"
	"github.com/alseRokachev/chat-app/internal/ws"
	"github.com/gin-gonic/gin"
	"net/http"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler) {
	r = gin.Default()

	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)

	r.POST("/ws/createRoom", wsHandler.CreateRoom)
}

func Start(ctx context.Context, addr string) error {
	server := &http.Server{
		Handler: r,
		Addr:    addr,
	}

	ch := make(chan error, 1)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			ch <- err
		}
		close(ch)
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		return server.Shutdown(ctx)
	}
}
