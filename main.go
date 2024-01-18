package main

import (
	"github.com/gin-gonic/gin"
)

func Main() {
	envport := "8000"

	router := gin.New()

	router.Run(":" + envport)
}
