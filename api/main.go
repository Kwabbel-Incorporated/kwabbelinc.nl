// main.go
package main

import (
	"github.com/Kwabbel-Incorporated/kwabbelinc.nl/database"
	"github.com/Kwabbel-Incorporated/kwabbelinc.nl/server"
	"github.com/Kwabbel-Incorporated/kwabbelinc.nl/util"
)

func main() {
    util.LoadEnvVariables()
    database.ConnectDatabase()
    server.StartServer()
}
