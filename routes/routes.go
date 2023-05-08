// Package routes represents package of API route definitions
package routes

import (
	"flight-tracking/handler"

	"github.com/gin-gonic/gin"
)

func CreateRouter(handler *handler.Handler) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	// Auto recover incase any unexpcted 500
	r.Use(gin.Recovery())

	r.POST("/calculate", handler.CalculatePath)

	return r
}
