package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	gin_handlers "gin-pg-login/internal/gin-handlers"
	"gin-pg-login/internal/server"
	"gin-pg-login/src/types"

	"github.com/joho/godotenv"
)

func main() {
	//==========================================================================
	// CONFIG
	//==========================================================================

	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err.Error())
	}

	//==========================================================================
	// DATABASE
	//==========================================================================

	database := types.NewDatabase()
	err = database.InitTables()
	if err != nil {
		log.Fatal(err.Error())
	}

	//==========================================================================
	// ROUTER
	//==========================================================================

	servHandlers := gin_handlers.NewGinHandlers(database, os.Getenv("DOMAIN"))

	routs := server.NewServo(servHandlers, os.Getenv("PORTN"))
	routs.InitRoutes()

	//==========================================================================
	// RUNNING
	//==========================================================================
	go func() {
		routs.Run()
	}()
	// Gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	<-quit
	log.Println("Shutting down server...")

	routs.KillServer()
	database.KillDB()
}
