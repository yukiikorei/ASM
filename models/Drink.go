/**
 * @Author: korei
 * @Description:
 * @File:  Drink
 * @Version: 1.0.0
 * @Date: 2020/11/17 上午10:51
 */

package models

type Drink struct {
	No			uint32 		`gorm:"primaryKey"`
	Name 		string 		`gorm:"type:varchar(30)"`
	Count       uint32
	Price 		float32
}
