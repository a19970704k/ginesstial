package main

import (
	"github.com/gin-gonic/gin"
	"lzh.practice/ginessential/controller"
	"lzh.practice/ginessential/midlleware"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.Use(midlleware.CORSMiddleware(), midlleware.RecoveryMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	//中间件保护路由
	r.GET("/api/auth/info", midlleware.AuthMiddleware(), controller.Info)

	categoryRouters := r.Group("/categories")
	categoryController := controller.NewCategoryController()
	categoryRouters.POST("", categoryController.Create)
	//put用一个model替换另一个model patch修改一部分
	categoryRouters.PUT("/:id", categoryController.Update)
	categoryRouters.GET("/:id", categoryController.Show)
	categoryRouters.DELETE("/:id", categoryController.Delete)

	postRouters := r.Group("/posts")
	postRouters.Use(midlleware.AuthMiddleware())
	postController := controller.NewPostController()
	postRouters.POST("", postController.Create)
	//put用一个model替换另一个model patch修改一部分
	postRouters.PUT("/:id", postController.Update)
	postRouters.GET("/:id", postController.Show)
	postRouters.DELETE("/:id", postController.Delete)
	postRouters.POST("/page/list", postController.PageList)

	return r
}
