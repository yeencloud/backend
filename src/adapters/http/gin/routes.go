package gin

import (
	"back/src/core/domain"
	"github.com/gin-gonic/gin"
)

func (server *ServiceHTTPServer) SetRoutes() {
	r := server.engine

	r.Use(server.getUserMiddleware())
	r.Use(server.getLangMiddleware())

	g := r.Use(server.verifySetupMiddleware())

	g.GET("/status", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "OK",
		})
	})

	r.NoRoute(func(context *gin.Context) {
		server.abortWithError(context, domain.ErrorNotFound)
	})

	r.NoMethod(func(context *gin.Context) {
		server.abortWithError(context, domain.ErrorNoMethod)
	})
}
