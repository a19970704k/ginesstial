package midlleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"lzh.practice/ginessential/common"
	"lzh.practice/ginessential/model"
)

//认证中间件
//可以对那些需要授权才能访问的接口进行验证。
//登录认证 权限校验 数据分页 记录日志
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")
		//validate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		//有效部分
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		//验证通过后获取token中claims里的userID
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		//查询userID并保存到user中
		DB.First(&user, userId)
		// 用户存不存在
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		//用户存在 将user信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}
