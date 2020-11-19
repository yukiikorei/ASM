/**
 * @Author: korei
 * @Description:
 * @File:  hardware
 * @Version: 1.0.0
 * @Date: 2020/11/18 下午4:47
 */

package api

var DrinkBoxInstance DrinkBox

type DrinkBox interface {
	Open(no uint32)
	OpenAndHeat(no uint32)
	OpenAndCool(no uint32)
}



