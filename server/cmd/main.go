package main

import(
	"log"
	"server/db"
	"server/internal/user"
	"server/internal/ws"
	"server/router"
)


func main() {
	dbConn,err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}

	userReq := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userReq)
	userHandler := user.NewHandler(userSvc)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouter(userHandler, wsHandler)
	router.Start("0.0.0.0:8080")
}