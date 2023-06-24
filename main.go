package main

import (
	"context"
	"htmx-scorecard/src/components"
	"htmx-scorecard/src/pages"
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

	pages.PageLogin(r, mongoStore)
	pages.PageSignup(r, mongoStore)

	//-----------------------------
	// COMPONENTS
	//-----------------------------

	components.NavGuestOpened(r)
	components.NavGuestClosed(r)

	//-----------------------------
	// SERVING
	//-----------------------------

	r.Run()

}