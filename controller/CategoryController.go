package controller

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"lzh.practice/ginessential/model"
	"lzh.practice/ginessential/repository"
	"lzh.practice/ginessential/response"
	"lzh.practice/ginessential/vo"
)

//为了使方法名可以服用 使用结构体方法（接收者）
type ICategoryController interface {
	RestController
}

type CategoryController struct {
	Repository repository.CategoryRepositpry
}

func NewCategoryController() ICategoryController {
	repository := repository.NewCategoryRepository()
	//自动创建表单
	repository.DB.AutoMigrate(model.Category{})
	return CategoryController{
		Repository: repository,
	}
}

func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory vo.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		log.Println(err.Error())
		response.Fail(ctx, nil, "数据验证错误，分类名称必填")
		return
	}
	category, err := c.Repository.Create(requestCategory.Name)
	if err != nil {
		panic(err)
		// return
	}
	response.Success(ctx, gin.H{"category": category}, "")
}

//要接受两方面的参数 body 和 path里的
func (c CategoryController) Update(ctx *gin.Context) {
	//绑定body中的参数
	var requestCategory vo.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		log.Println(err.Error())
		response.Fail(ctx, nil, "数据验证错误，分类名称必填")
		return
	}
	//获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	updateCategory, err := c.Repository.SelectById(categoryId)
	if err != nil {
		panic(err)
	}
	//更新
	//map、struct、name value
	category, err := c.Repository.Update(*updateCategory, requestCategory.Name)
	if err != nil {
		panic(err)
	}
	response.Success(ctx, gin.H{"category": category}, "修改成功")
}

func (c CategoryController) Show(ctx *gin.Context) {
	//获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	category, err := c.Repository.SelectById(categoryId)
	if err != nil {
		panic(err)
	}
	response.Success(ctx, gin.H{"category": category}, "展示成功")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	//获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	var category model.Category
	if err := c.Repository.DeletById(categoryId); err != nil {
		response.Fail(ctx, nil, "删除失败请重试")
		return
	}
	response.Success(ctx, gin.H{"category": category}, "删除成功")
}
