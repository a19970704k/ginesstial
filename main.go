package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"lzh.practice/ginessential/common"

	//一定要记得插入驱动
	_ "github.com/go-sql-driver/mysql"
)

//使用gorm使用数据库

func main() {
	//读取配置
	InitConfig()
	//连接数据库
	db := common.InitDB()
	//延迟关闭
	defer db.Close()
	r := gin.Default()
	r = CollectRouter(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run()) // listen and serve on 0.0.0.0:8080
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
