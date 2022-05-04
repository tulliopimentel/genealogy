package main

import (
	"example/desafio/database"
	"example/desafio/server"
)

// @title           			Genealogy API
// @version         			1.0
// @description     			This is the documentation of the Genealogy API
// @termsOfService  			http://swagger.io/terms/
// @tulliopimentelb@gmail.com   API Support
// @contact.url    				https://www.linkedin.com/in/tulliopimentelbarbosa/
// @license.name  				Apache 2.0
// @license.url   				http://www.apache.org/licenses/LICENSE-2.0.html
// @host      					localhost:5050
func main() {

	database.StartDb()

	server := server.NewServer()

	server.Run()
}
