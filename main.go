/**
 * @Author: korei
 * @Description:
 * @File:  main.go
 * @Version: 1.0.0
 * @Date: 2020/11/17 上午1:05
 */

package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
	"sync"

	"korei/ASM/datasource"
	"korei/ASM/services"
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
	app.Use(logger.New())

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

func main1() {
	datasource.InitDB()

	//m := services.MachineServices{}
	s := services.SellingService{}

	/*
	d1 := models.Drink{
		No: 1,
		Name: "drink1",
		Price: 3.5,
		Count: 10,
	}

	d2 := models.Drink{
		No: 2,
		Name: "drink_hhc",
		Price: 3,
		Count: 8,
	}

	m.AddDrink(&d1)
	m.AddDrink(&d2)
	 */

	err := s.Sell(1,"hot")
	if err!= nil {
		fmt.Println("1. ",err)
	}
	err =s.Sell(2,"normal")
	if err!= nil {
		fmt.Println("2. ",err)
	}


	//wait()
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()

}
