package main

import (
	"github.com/kooroshh/fiber-boostrap/bootstrap"
	"github.com/kooroshh/fiber-boostrap/router"
	"log"
)

func main() {
	app := bootstrap.NewApplication()
	router.InstallRouter(app)
	log.Fatal(app.Listen(":3000"))
}
