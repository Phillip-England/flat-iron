package main

import (
	"context"
	"htmx-scorecard/src/routes/actionroutes"
	"htmx-scorecard/src/routes/componentroutes"
	"htmx-scorecard/src/routes/pageroutes"
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

	pageroutes.Login(r, mongoStore)
	pageroutes.Signup(r, mongoStore)
	pageroutes.LocationSelection(r, mongoStore)
	pageroutes.ServerError(r)
	pageroutes.Unauthorized(r)
	pageroutes.Test(r, mongoStore)
	pageroutes.Loading(r, mongoStore)

	//-----------------------------
	// COMPONENTS
	//-----------------------------
	
	componentroutes.NavGuestOpened(r)
	componentroutes.NavGuestClosed(r)
	componentroutes.NavUserOpened(r)
	componentroutes.NavUserClosed(r)

	//-----------------------------
	// SERVER ACTIONS
	//-----------------------------

	actionroutes.LoginUser(r, mongoStore)
	actionroutes.SignupUser(r, mongoStore)
	actionroutes.CreateLocation(r, mongoStore)
	actionroutes.Logout(r)
	
	//-----------------------------
	// SERVING
	//-----------------------------

	r.Run()

}