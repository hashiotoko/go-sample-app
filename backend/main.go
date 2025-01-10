package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashiotoko/go-sample-app/backend/config"
	"github.com/hashiotoko/go-sample-app/backend/infrastructure"
	"github.com/hashiotoko/go-sample-app/backend/infrastructure/db"
	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadConfig()
	// fmt.Printf("config: %+v\n", config.Config)
	cj, err := json.MarshalIndent(config.Config, "", " ")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println("config:", string(cj))

	infrastructure.InitLogger()

	router := echo.New()
	dbClient := db.NewClient()
	infrastructure.Init(router, dbClient)

	router.Start(":8888")
}
