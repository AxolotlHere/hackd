package main

import (
	"fmt"
	"log"
)

func main() {
	initDbController()
	log.Println(createUser("him", "im him again", "againimhim@gmail.com", "GoB!g123"))
	app := &app{
		server_config: config{
			Addr: ":8080",
		},
	}
	router_mux := app.mount()
	fmt.Println(app.initialiseServer(router_mux))
}
