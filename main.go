package main

import (
	"fmt"
	"github.com/shirou/gopsutil/mem"
	"log"
	"math"
	"strconv"
	"time"
	"github.com/shirou/gopsutil/cpu"
	"github.com/getlantern/systray"

)

func main() {
	systray.Run(onReady, onExit)

}

func onReady() {
	model,numberOfCore,frequency,cacheSize:=getInfo()
	go func() {
		var result string
		for {
			result = getData()
			systray.SetTitle(result)
		}

	}()
	systray.AddMenuItem(fmt.Sprintf("CPU            : %s",model),"Cpu Model")
	systray.AddMenuItem(fmt.Sprintf("Cores          : %74s",strconv.Itoa(int(numberOfCore))),"Number of core")
	systray.AddMenuItem(fmt.Sprintf("Frequency  : %70s",strconv.Itoa(frequency)),"Frequency CPU")
	systray.AddMenuItem(fmt.Sprintf("CPU Cache : %71s",strconv.Itoa(int(cacheSize))),"CPU Cache")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quits this app")
	go func() {
		for {
			select {
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}

			}

	}()
}

func onExit(){

}

func getMemoryUsage()int{
	memory,err:=mem.VirtualMemory()
	if err!=nil{
		log.Fatal(err)
	}
	return int(math.Ceil(memory.UsedPercent))
}

func getCpuUsage()int{
	percent,err:=cpu.Percent(time.Second,false)
	if err!=nil{
		log.Fatal(err)
	}
	return int(math.Ceil(percent[0]))
}

func getData()(string){
	cpuData:="Cpu: "+strconv.Itoa(getCpuUsage())+"% "
	memoryData:="Mem: "+strconv.Itoa(getMemoryUsage())+"% "
	return cpuData+memoryData
}

func getInfo()(string,int32,int,int32){
	info,err:=cpu.Info()
	if err!=nil{
		log.Fatal(err)
	}
	return info[0].ModelName,info[0].Cores,int(info[0].Mhz),info[0].CacheSize
}