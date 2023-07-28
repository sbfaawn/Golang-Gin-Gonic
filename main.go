package main

import (
	"Golang-Gin-Gonic/router"
	"fmt"
)

func main() {
	fmt.Println("Gin Gonic Server is Running")

	r := router.NewRouter()
	r.Run(":8085") // listen and serve on 0.0.0.0:8080
}
