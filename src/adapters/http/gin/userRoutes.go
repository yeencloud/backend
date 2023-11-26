package gin

import "github.com/gin-gonic/gin"

func (server *ServiceHTTPServer) getUserMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (server *ServiceHTTPServer) getLangMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		_, exists := context.Get("user")
		if !exists {
			acceptLanguage := context.GetHeader("Accept-Language")
			if acceptLanguage != "" {
				context.Set("lang", acceptLanguage)
				return
			}
			context.Set("lang", "en")
		}
	}
}
