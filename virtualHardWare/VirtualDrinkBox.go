/**
 * @Author: korei
 * @Description:
 * @File:  VirtualDrinkBox
 * @Version: 1.0.0
 * @Date: 2020/11/18 下午6:25
 */

package virtualHardWare

import (
	"fmt"
	"korei/ASM/api"
	"sync"
	"time"
)

type VirtualDrinkBox struct {

}

func (v *VirtualDrinkBox) Open(no uint32)  {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		wg.Done()
		api.BeginWork()
		defer api.EndWork()
		fmt.Printf("open %v \n ", no)
		time.Sleep(time.Second*5)
		fmt.Printf("finish \n")
	}()
	wg.Wait()
}

func (v *VirtualDrinkBox) OpenAndHeat(no uint32) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		wg.Done()
		api.BeginWork()
		defer api.EndWork()
		fmt.Printf("open %v \n ", no)
		time.Sleep(time.Second*5)
		fmt.Printf("heating \n")
		time.Sleep(time.Second*10)
		fmt.Printf("finish \n")
	}()
	wg.Wait()
}

func (v *VirtualDrinkBox) OpenAndCool(no uint32) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		wg.Done()
		api.BeginWork()
		defer api.EndWork()
		fmt.Printf("open %v \n ", no)
		time.Sleep(time.Second*5)
		fmt.Printf("cooling \n")
		time.Sleep(time.Second*10)
		fmt.Printf("finish \n")
	}()
	wg.Wait()
}


