package urlshorten

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var redis = RedisConn{}

// todo: env var this
const SHORTPREFIX = "http://localhost:10000/api/v1/"

type tempURL struct {
	TempURL string `json:"url"`
}

func RedirectToLong(c *gin.Context) {
	log.Println("endpoint reached: redirect to long url")

	shortURL := c.Params.ByName("url")
	redirectURL := redis.RetrieveURL(shortURL)

	log.Printf("redirecting to: %s\n", redirectURL)
	c.Redirect(http.StatusTemporaryRedirect, redirectURL)
}

func GenerateShort(c *gin.Context) {
	log.Println("endpoint reached: generate short url")

	var holder tempURL
	c.ShouldBind(&holder)

	shortURL, err := ShortenURL(holder.TempURL)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("unexpected error"))
	}
	_ = redis.PublishURL(shortURL, holder.TempURL)

	log.Println("returning shortened url")
	c.JSON(http.StatusOK, map[string]string{"short_url": SHORTPREFIX + shortURL})
}
