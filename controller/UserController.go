package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"lzh.practice/ginessential/common"
	"lzh.practice/ginessential/dto"
	"lzh.practice/ginessential/model"
	"lzh.practice/ginessential/response"
	"lzh.practice/ginessential/utils"
)

func Register(ctx *gin.Context) {
	//获取数据库
	db := common.GetDB()
	//获取参数
	//使用map获取请求参数
	// var requestMap = make(map[string]string)
	// json.NewDecoder(ctx.Request.Body).Decode(&requestMap)
	//使用struct获取请求参数
	var requestUser = model.User{}
	// json.NewDecoder(ctx.Request.Body).Decode(&requestUser)
	ctx.Bind(&requestUser)
	//gin的bind函数
	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password
	//数据验证
	if len(telephone) != 11 {
		log.Println("....", telephone)
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}

	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}

	if len(name) == 0 {
		name = utils.RandomString(10)
	}
	log.Println(name, telephone, password)
	//判断手机号是否存在
	if isTelephoneExit(db, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已存在不允许注册")
		return
	}
	//创建用户
	//上述判断用户不存在 可以创建新用户
	//golang内置的bcrypt加密算法 对password进行加密 用户密码不能明文保存
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	db.Create(&newUser)

	//发放token给前端
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统错误"})
		log.Printf("token generate error : %v", err)
		return
	}
	//返回结果
	//gin.H --  map[string]interface{}
	response.Success(ctx, gin.H{"token": token}, "注册成功")

}

func Login(ctx *gin.Context) {
	//引入DB
	db := common.GetDB()
	//获取参数
	//使用struct获取请求参数
	var requestUser = model.User{}
	// json.NewDecoder(ctx.Request.Body).Decode(&requestUser)
	ctx.Bind(&requestUser)
	//gin的bind函数
	// name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password
	//数据验证
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}

	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
		return
	}
	//手机号是否存在
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		return
	}
	//发放token给前端
	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统错误"})
		log.Printf("token generate error : %v", err)
		return
	}
	//返回结果
	//gin.H --  map[string]interface{}
	response.Success(ctx, gin.H{"token": token}, "登录成功")
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}})
}

func isTelephoneExit(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	return user.ID != 0
}
