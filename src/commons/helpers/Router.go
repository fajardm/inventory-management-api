package helpers

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func SetRouter() *gin.Engine {
	router = gin.Default()
	return router
}

func GetRouter() *gin.Engine {
	return router
}
