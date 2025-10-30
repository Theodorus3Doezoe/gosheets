package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
	"repsheets-go/routes"
    // mappen
    "repsheets-go/database"
)

func main() {
    err := database.Connect()
    if err != nil {
        panic("Database verbinding mislukt: " + err.Error())
    }

    fmt.Println("Database verbinding succesvol!")

    fmt.Println("Router wordt opgezet...")
    r := gin.Default()

	routes.SetupRoutes(r, database.GetDB())

    fmt.Println("Server start op http://localhost:8080")
    r.Run(":8080")
}
