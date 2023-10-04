package main

import (
	"log"

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

	router.InitRouter(userHandler)
	err = router.Start("0.0.0.0:8080")
	if err != nil {
		log.Fatal(err)
	}
}
