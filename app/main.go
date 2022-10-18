package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const defaultEnvValue = "[ENV VAR NOT SET]"

func getEnvVar(key string) string {
	value := os.Getenv(key)

	if value == "" {
		value = defaultEnvValue
	}

	return value
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	pingMessage := fmt.Sprintf("pong! Company Name: %v", getEnvVar("DEVOPSTESTE_NAME"))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": pingMessage,
		})
	})

	httpPort := getEnvVar("DEVOPSTESTE_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	router.Run(":" + httpPort)
}
