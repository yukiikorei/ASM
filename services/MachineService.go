/**
 * @Author: korei
 * @Description:
 * @File:  MachineService
 * @Version: 1.0.0
 * @Date: 2020/11/17 下午2:04
 */

package services

import (
	"korei/ASM/api"
	"korei/ASM/datasource"
	"korei/ASM/models"
	"sync"
)

var (
	isOpen = false
	openMutex sync.Mutex
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

func (M *MachineServices) IsOpen() bool {
	openMutex.Lock()
	defer openMutex.Unlock()

	return isOpen
}

func (M *MachineServices) TryOpen() bool {
	openMutex.Lock()
	defer openMutex.Unlock()

	if isOpen {
		return true
	}
	select {
	case <-api.Work:
		isOpen = true
		return true
	default:
		return false
	}
}

func (M *MachineServices) Close() bool {
	openMutex.Lock()
	defer openMutex.Unlock()

	if !isOpen {
		return false
	}else {
		isOpen = false
		api.Work<-"close"
		return true
	}
}

