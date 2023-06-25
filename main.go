package main

import (
	"context"
	"htmx-scorecard/src/routes"
	"htmx-scorecard/src/types"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	//-----------------------------
	// CONFIG
	//-----------------------------

	godotenv.Load()
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	r.Static("/static", "./static")

	//-----------------------------
	// DATABASE
	//-----------------------------

	mongoStore, err := types.NewMongoStore()
	if err != nil {
		log.Fatal("failed to connect to mongo db")
	}
	defer mongoStore.Client.Disconnect(context.Background())

	//-----------------------------
	// PAGES
	//-----------------------------

	routes.PageLogin(r, mongoStore)
	routes.PageSignup(r, mongoStore)
	routes.PageLocationSelection(r, mongoStore)
	routes.PageServerError(r)

	//-----------------------------
	// COMPONENTS
	//-----------------------------
	
	routes.NavGuestOpened(r)
	routes.NavGuestClosed(r)
	routes.NavUserOpened(r)
	routes.NavUserClosed(r)

	//-----------------------------
	// SERVER ACTIONS
	//-----------------------------

	routes.ActionLoginUser(r, mongoStore)
	routes.ActionCreateLocation(r, mongoStore)

	//-----------------------------
	// SERVING
	//-----------------------------

	r.Run()

}