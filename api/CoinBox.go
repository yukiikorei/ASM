/**
 * @Author: korei
 * @Description:
 * @File:  CoinBox.go
 * @Version: 1.0.0
 * @Date: 2020/11/18 下午6:47
 */

package api

var CoinBoxInstance CoinBox

type CoinBox interface {
	Cost(count float32)bool
}
