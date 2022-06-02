package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var port string

func main() {
	mux := gin.Default()

	mux.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Status": "OK",
		})
	})

	if err := mux.Run(port); err != nil {
		log.Fatal(err.Error())
	}

}

func init() {
	if err := godotenv.Load(); err != nil {
		new_port, exists := os.LookupEnv("PORT")
		if exists {
			port = new_port
		} else {
			log.Fatal(err.Error(), ": no port")
			return
		}
	} else {
		port = os.Getenv("PORT")
	}
	//	db := sqlx.Open("postgres", fmt.Sprintf())
}
