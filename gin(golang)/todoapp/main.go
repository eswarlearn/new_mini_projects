package main

import (
	"codesignal.com/todoapp/router"
)

func main() {
	r := router.SetupRoute()
	r.Run(":8080")
}
