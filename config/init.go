/**
 * @Author: korei
 * @Description:
 * @File:  init.go
 * @Version: 1.0.0
 * @Date: 2020/11/17 下午12:05
 */

package config

import "github.com/BurntSushi/toml"

var GlobalConfig Config

type Config struct{
	DBConfig 		DBConfig
	AdminConfig 	AdminConfig
	MachineConfig 	MachineConfig
}

type AdminConfig struct {
	Pass 			string
	Name 			string
}

type MachineConfig struct {
	KindOfDrink		uint32
}

type DBConfig struct{
	User 	string
	Pass 	string
	Host 	string
	Port 	string
	Name 	string
}

func init_local() {
	// init db
	GlobalConfig.DBConfig.User = "root"
	GlobalConfig.DBConfig.Pass = "test"
	GlobalConfig.DBConfig.Host = "127.0.0.1"
	GlobalConfig.DBConfig.Port = "3306"
	GlobalConfig.DBConfig.Name = "ASMDB"

	//init machine
	GlobalConfig.MachineConfig.KindOfDrink = 16

	//admin config
	GlobalConfig.AdminConfig.Name = "admin"
	GlobalConfig.AdminConfig.Pass = "password"
}

func init(){
	configPath := "./config/config"
	_,err := toml.DecodeFile(configPath,&GlobalConfig)
	if err!= nil {
		panic(err)
	}
}