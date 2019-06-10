package routers

import (
	"adminVideos/routers/api"
	"adminVideos/routers/api/v1"
	"adminVideos/routers/api/v3"
	"github.com/gin-gonic/gin"
)
/*
val,_:=c.Get("persons_id")
	fmt.Println("val:",val)
*/
func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	//静态文件服务 提供图片路径
	r.Static("/images","./public/upload/images")
	r.Static("/static/images","./public/images")
	r.Static("/videos","./public/upload/videos")
	r.Static("/index.html","./public/dist")
	
	//登录模块
	api1 := r.Group("/api/v1")
	{
		api1.POST("/login",v1.Login)
		api1.POST("/registered",v1.Registered)
	}

	//用户管理模块
	api2 := r.Group("/api/v2")
	api2.Use(api.GetAuth)
	{

	}
	api3 := r.Group("/api/v3")
	api3.Use(api.GetAuth)
	{
		//上传接口
		api3.POST("/upload/images",v3.UploadImages) //上传视频封面
		api3.POST("/upload/videos",v3.UploadVideos) //上传视频

		api3.POST("/class",v3.Class)  //创建菜单

		api3.POST("/video",v3.PostVideo)  //创建视频 创建和查询使用Raw和Raws

		api3.PUT("/video/:id",v3.UpdateVideo)  //更新视频 更新的时候使用Db.Exec方法
		api3.GET("/video",v3.QueryVideo)

	}





return r
}
