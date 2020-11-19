/**
 * @Author: korei
 * @Description:
 * @File:  status.go
 * @Version: 1.0.0
 * @Date: 2020/11/18 下午7:27
 */

package api

var Work chan string

func init()  {
	Work = make(chan string,1)
	Work <- "begin"
}


func BeginWork()  {
	<-Work
}

func EndWork()  {
	Work<-"end"
}