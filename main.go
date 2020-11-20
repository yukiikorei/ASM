/**
 * @Author: korei
 * @Description:
 * @File:  main.go
 * @Version: 1.0.0
 * @Date: 2020/11/17 上午1:05
 */

package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
	"korei/ASM/datasource"
	_ "korei/ASM/virtualHardWare"
	"korei/ASM/web/controllers"
)

func main()  {

	datasource.InitDB()

	//1. create iris servers
	app := iris.New()
	//2. set recovery
	app.Use(recover.New())
	//3. set logger
	//app.Use(logger.New())

	//. set error process

	//. set mvc handler
	mvc.New(app.Party("/sell")).Handle(new(controllers.SellingController))
	mvc.New(app.Party("/admin")).Handle(new(controllers.AdminController))


	//. register view and fs
	app.RegisterView(iris.HTML("./web/views/",".html"))
	app.HandleDir("/content","./web/content")

	//. run the app
	app.Run(iris.Addr(":10101"))


}
