package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Ahmed-Mas/url_shorten_back/urlshorten"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}

	r := gin.Default()

	r.POST("/api/v1/short", urlshorten.GenerateShort)
	r.GET("/api/v1/:url", urlshorten.RedirectToLong)

	log.Println("starting server at")
	r.Run(fmt.Sprintf(":%s", port))
}
