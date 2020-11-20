/**
 * @Author: korei
 * @Description:
 * @File:  AdminController.go
 * @Version: 1.0.0
 * @Date: 2020/11/17 上午1:19
 */

package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/basicauth"
	"github.com/kataras/iris/v12/mvc"
	"korei/ASM/config"
	"korei/ASM/services"
	"strconv"
)

type AdminController struct {
	Ctx iris.Context
}

// set middleware
func (c *AdminController) BeforeActivation(b mvc.BeforeActivation)  {
	//1. add basic auth middleware
	authConfig := basicauth.Config{
		Users: map[string]string{
			config.GlobalConfig.AdminConfig.Name : config.GlobalConfig.AdminConfig.Pass,
		},
		Realm: "auth require",
		Expires: 0,
	}
	authentication := basicauth.New(authConfig)
	b.Router().Use(authentication)

	//2. open machine middleware
	b.Router().Use(func(ctx iris.Context) {
		m := services.MachineServices{}
		if m.IsOpen() || ctx.Path()=="/admin/open"{
			ctx.Next()
		}else{
			ctx.View("open.html",iris.Map{"content":""})
		}
	})

}

func (c *AdminController) Get() (result mvc.Result){

	m := services.MachineServices{}
	drinks,_ := m.GetDrinkList()

	return mvc.View{
		Name : "admin.html",
		Data : iris.Map{
			"drinks" : drinks,
		},
		Layout: "admin.html",
	}
}

func (c *AdminController) GetOpen() (result mvc.Result){
	m := services.MachineServices{}
	suc := m.TryOpen()
	if suc {
		c.Ctx.Redirect("/admin")
		return
	}else{
		return mvc.View{
			Name: "/admin/open.html",
			Data: iris.Map{
				"content" : "machine is working",
			},
			Layout: "open.html",
		}
	}
}

func (c *AdminController) GetClose()  {
	m := services.MachineServices{}
	m.Close()
	c.Ctx.Redirect("/admin")
}

func (c *AdminController) PostUpdate()  {
	temNo,_ := strconv.ParseUint( c.Ctx.URLParam("No"),10,32)
	No := uint32(temNo)
	Attr := c.Ctx.URLParam("Attr")
	value := c.Ctx.URLParam("value")
	m := services.MachineServices{}

	if Attr=="Name" {
		m.UpdateName(No,value)
	}else if Attr=="Count"{
		temCount,_ := strconv.ParseUint(value,10,32)
		count := uint32(temCount)
		m.UpdateCount(No,count)
	}else if Attr=="Price" {
		temPrice,_ := strconv.ParseFloat(value,32)
		price := float32(temPrice)
		m.UpdatePrice(No,price)
	}
}

func (c *AdminController) PostUpload() (result mvc.Result){
	return c.Get()
}
