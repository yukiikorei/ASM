/**
 * @Author: korei
 * @Description:
 * @File:  MachineService
 * @Version: 1.0.0
 * @Date: 2020/11/17 下午2:04
 */

package services

import (
	"korei/ASM/datasource"
	"korei/ASM/models"
)

type MachineServices struct {

}

/*
func (M* MachineServices)AddDrink(drink *models.Drink) (err error){
	err = datasource.DB.Create(drink).Error
	return
}
 */

/*
func (M* MachineServices)DeleteDrink(drink *models.Drink)(err error) {
	err = datasource.DB.Unscoped().Delete(drink).Error
	return
}
 */

func (M* MachineServices)GetDrinkList() (drinks []models.Drink,err error) {
	err = datasource.DB.Find(&drinks).Error
	return
}

func (M* MachineServices)UpdateDrink(drink *models.Drink)(err error){
	err = datasource.DB.Updates(drink).Error
	return
}


func (M* MachineServices)UpdatePrice(No uint32,price float32)(err error){
	err = datasource.DB.Model(&models.Drink{}).Where("no=?",No).Update("price",price).Error
	return
}

func (M *MachineServices) UpdateCount(No uint32,Count uint32) (err error) {
	err = datasource.DB.Model(&models.Drink{}).Where("no=?",No).Update("count",Count).Error
	return
}

func (M *MachineServices) UpdateName(No uint32,Name string) (err error) {
	err = datasource.DB.Model(&models.Drink{}).Where("no=?",No).Update("name",Name).Error
	return
}