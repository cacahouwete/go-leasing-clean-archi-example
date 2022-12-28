package main

import (
	"log"

	"gitlab.com/alexandrevinet/leasing/app"

	"gitlab.com/alexandrevinet/leasing/config"
)

// @title       Leasing
// @version     1.0
// @description An api to manage leasing locations.
//
// @contact.name   Alexandre VINET
// @contact.email  contact@alexandrevinet.fr
//
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
//
// @BasePath    /api/v1
//
// main.
func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(*cfg)
}
