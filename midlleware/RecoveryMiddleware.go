package midlleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"lzh.practice/ginessential/response"
)

//err信息输出到前台 要对panic进行拦截
func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				response.Fail(ctx, nil, fmt.Sprint(err))
			}
		}()
		ctx.Next()
	}
}
