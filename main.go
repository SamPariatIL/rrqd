package main

import (
	"github.com/SamPariatIL/rrqd/cmd"
	"github.com/SamPariatIL/rrqd/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.SetupRoutes(router)

	cmd.Execute()

	router.Run(":8080")
}
