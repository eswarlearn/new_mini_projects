package main

import (
	"github.com/eswarlearn/go-sms-verify/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//initialize config
	app := api.Config{Router: router}

	app.Routes()

	router.Run(":8081") // Start the server on port 8081
}
