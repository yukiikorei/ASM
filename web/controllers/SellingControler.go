/**
 * @Author: korei
 * @Description:
 * @File:  sellingControler
 * @Version: 1.0.0
 * @Date: 2020/11/17 上午1:18
 */

package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"korei/ASM/api"
	"korei/ASM/services"
	"strconv"
)

type SellingController struct {
	Ctx iris.Context
}

func (c *SellingController) Get() (result mvc.Result){
	select {
	case <-api.Work:
		s := services.SellingService{}
		drinks,_ := s.GetDrinks()
		result = mvc.View{
			Name:"/sell/index.html",
			Data: iris.Map{
				"drinks" : drinks,
			},
			Layout: "index.html",
		}
		api.Work <- " "
	default:
		result = mvc.View{
			Name:"/sell/working.html",
			Data: iris.Map{},
			Layout: "working.html",
		}
	}

	return
}

func (c *SellingController) GetBuy() (result mvc.Result){
	noString 	:= c.Ctx.URLParam("No")
	temp:= c.Ctx.URLParam("temp")
	no,_ := strconv.ParseUint(noString,10,32)
	s := services.SellingService{}

	err := s.Sell(uint32(no),temp)
	if err==nil {
		c.Ctx.Redirect("/sell")
	}else{
		result = mvc.View{
			Name: "error.html",
			Data: iris.Map{
				"msg" : err.Error(),
			},
			Layout: "sellError.html",
		}
	}

	return
}


