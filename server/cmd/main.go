package main

import (
	"context"
	"github.com/alseRokachev/chat-app/internal/ws"
	"log"
	"os"
	"os/signal"

	"github.com/alseRokachev/chat-app/db"
	"github.com/alseRokachev/chat-app/internal/user"
	"github.com/alseRokachev/chat-app/router"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatal("could not init db conn: ", err)
	}

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)

	go hub.Run()

	router.InitRouter(userHandler, wsHandler)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err = router.Start(ctx, "0.0.0.0:8001")
	if err != nil {
		log.Fatal(err)
	}
}
