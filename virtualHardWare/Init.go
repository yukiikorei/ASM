/**
 * @Author: korei
 * @Description:
 * @File:  Init
 * @Version: 1.0.0
 * @Date: 2020/11/18 下午6:47
 */

package virtualHardWare

import "korei/ASM/api"

func init() {
	var vcb VirtualCoinBox
	vcb.Work()

	api.DrinkBoxInstance = new(VirtualDrinkBox)
	api.CoinBoxInstance  = &vcb
}