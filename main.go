package main

import (
	"context"
	"htmx-scorecard/src/routes/actionRoutes"
	"htmx-scorecard/src/routes/componentRoutes"
	"htmx-scorecard/src/routes/pageRoutes"
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

	pageRoutes.Login(r, mongoStore)
	pageRoutes.Signup(r, mongoStore)
	pageRoutes.LocationSelection(r, mongoStore)
	pageRoutes.ServerError(r)
	pageRoutes.Unauthorized(r)

	//-----------------------------
	// COMPONENTS
	//-----------------------------
	
	componentRoutes.NavGuestOpened(r)
	componentRoutes.NavGuestClosed(r)
	componentRoutes.NavUserOpened(r)
	componentRoutes.NavUserClosed(r)

	//-----------------------------
	// SERVER ACTIONS
	//-----------------------------

	actionRoutes.LoginUser(r, mongoStore)
	actionRoutes.CreateLocation(r, mongoStore)
	actionRoutes.Logout(r)

	//-----------------------------
	// SERVING
	//-----------------------------

	r.Run()

}