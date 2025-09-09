package main

import (
	"fmt"
)

func main() {
	initDbController()
	app := &app{
		server_config: config{
			Addr: ":8080",
		},
	}
	router_mux := app.mount()
	fmt.Println(app.initialiseServer(router_mux))
}
