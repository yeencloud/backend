package gin

import (
	"github.com/gin-gonic/gin"
)

func (server *ServiceHTTPServer) verifySetupMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		isDone, nextStep := server.ucs.IsSetupDone()

		if isDone {
			context.Next()
		} else {
			_ = nextStep
			//server.abortWithError(context, 500, errors.New("setup not done: "+nextStep))
		}
	}
}
