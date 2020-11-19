/**
 * @Author: korei
 * @Description:
 * @File:  datasource.go
 * @Version: 1.0.0
 * @Date: 2020/11/17 上午11:19
 */

package datasource

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"korei/ASM/models"
	"sync"

	"korei/ASM/config"
)

var (
	DB             	* gorm.DB
	dbConnectMutex 	sync.Mutex

	/*
	Coins          	float32
	coinsMutex		sync.Mutex
	 */
)

//connect to database
func Connect()  {
	dbConnectMutex.Lock()
	defer dbConnectMutex.Unlock()

	dbc := config.GlobalConfig.DBConfig

	dsn := dbc.User+":"+dbc.Pass+"@tcp("+dbc.Host+":"+dbc.Port+")/"+
		dbc.Name + "?charset=utf8mb4&parseTime=True&loc=Local"

	temDB,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err!= nil {
		panic(err)
	}else{
		DB = temDB
	}
}

// create table
func InitDB() {
	if DB == nil {
		Connect()
	}
	mig := DB.Migrator()
	if !mig.HasTable(&models.Drink{}){
		err := mig.CreateTable(&models.Drink{})
		if err!= nil{
			panic(err)
		}
		kinds := config.GlobalConfig.MachineConfig.KindOfDrink
		for i:=uint32(1);i<=kinds;i++ {
			newDrink := models.Drink{
				No: i,
				Name: "",
				Count: 0,
				Price: 0,
			}
			DB.Create(&newDrink)
		}
	}
}

/*
func InitCoins()  {
	coinsMutex.Lock()
	defer coinsMutex.Unlock()
	Coins = 0
}

func AddCoins(count float32)  {
	coinsMutex.Lock()
	defer coinsMutex.Unlock()
	Coins += count
}

func GiveChange() {
	coinsMutex.Lock()
	defer coinsMutex.Unlock()
	Coins = 0
}
 */