package main

import (
	"fmt"
	"github.com/laioncorcino/personality/database"
	"github.com/laioncorcino/personality/route"
)

func main() {
	fmt.Println("Iniciando servidor Rest em Go ..")
	database.ConnectDatabase()
	route.HandleRequest()
}
