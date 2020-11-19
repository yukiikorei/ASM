/**
 * @Author: korei
 * @Description:
 * @File:  sellingService
 * @Version: 1.0.0
 * @Date: 2020/11/17 下午2:03
 */

package services

import (
	"errors"
	"korei/ASM/api"
	"korei/ASM/datasource"
	"korei/ASM/models"
)

type SellingService struct{

}

func (s* SellingService) GetDrinks() (drinks []models.Drink,err error) {
	err = datasource.DB.Where("count>0").Find(&drinks).Error
	return
}

func (s * SellingService)Sell(drinkNo uint32, temperature string)(err error){
	drink := models.Drink{
		No: drinkNo,
	}
	err = datasource.DB.Find(&drink).Error
	if err!= nil {
		return err
	}

	//check count
	if drink.Count==0 {
		return errors.New("sold out")
	}

	//check coin
	if !api.CoinBoxInstance.Cost(drink.Price) {
		return errors.New("money not enough")
	}
	//datasource.DB.Updates(&drink)
	datasource.DB.Model(&drink).Where("no = ?",drink.No).Update("count",drink.Count-1)

	if temperature=="hot"{
		api.DrinkBoxInstance.OpenAndHeat(drinkNo)
	}else if temperature=="cool"{
		api.DrinkBoxInstance.OpenAndCool(drinkNo)
	}else {
		api.DrinkBoxInstance.Open(drinkNo)
	}

	return nil
}

