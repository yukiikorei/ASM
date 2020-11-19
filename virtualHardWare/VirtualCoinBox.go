/**
 * @Author: korei
 * @Description:
 * @File:  VirtualCoinBox
 * @Version: 1.0.0
 * @Date: 2020/11/18 下午7:08
 */

package virtualHardWare

import (
	"fmt"
	"sync"
)

type VirtualCoinBox struct {
	count 			float32
	countMutex 		sync.Mutex
}

func (v *VirtualCoinBox) Cost(count float32) bool {
	v.countMutex.Lock()
	defer v.countMutex.Unlock()
	if count<=v.count {
		v.count-=count
		v.updateMonitor()
		return true
	}else {
		return false
	}
}

func (v *VirtualCoinBox) updateMonitor()  {
	fmt.Printf("monitor show : %v \n",v.count)
}

func (v *VirtualCoinBox) Work() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		wg.Done()
		var input uint32
		v.count = 0
		v.updateMonitor()
		fmt.Printf("Enter\n\t0 : change\n\t1 : enter 1¥\n\t5 : enter 0.5¥\n")
		for {
			fmt.Scan(&input)
			switch input {
			case 0:
				v.change()
			case 1:
				v.enter(1)
			case 5:
				v.enter(0.5)
			default:
				fmt.Printf("please input 0 / 1 / 5 / ^C \n")
			}
		}
	}()
	wg.Wait()
}

func (v *VirtualCoinBox) change()  {
	v.countMutex.Lock()
	defer v.countMutex.Unlock()
	fmt.Printf("change %v¥\n",v.count)
	v.count = 0
	v.updateMonitor()
}

func (v *VirtualCoinBox) enter(count float32)  {
	v.countMutex.Lock()
	defer v.countMutex.Unlock()
	v.count += count
	v.updateMonitor()
}
