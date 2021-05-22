package midlleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//允许访问的域名
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		//缓存时间
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		//可以通过访问的方法
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		//允许请求带的header信息
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}

		// ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		// ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		// ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		// ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		// ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	}
}
