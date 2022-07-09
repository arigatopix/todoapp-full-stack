package main

import (
	"server/models"
	r "server/routers"
)

func main() {
	models.ConnectDB()

	r.InitRouter()
}
