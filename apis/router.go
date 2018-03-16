package apis

import "github.com/gin-gonic/gin"

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func NewRouter() *gin.Engine {
	return router
}
