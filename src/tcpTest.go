package main

import "runtime"
import "comsManager"
import "elevNet"

import "fmt"
//import "strings"



const TARGET_IP = "129.241.187.142"



func main() {
	elevNet.TcpChanInit()
	elevNet.ComsChanInit()
	comsManager.NetChanInit()
	c := make(chan int)
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("hai")
	
	 
	go elevNet.HandleTCPCom()
	go elevNet.ListenToBroadcast(elevNet.ComsChan)
	go comsManager.DeliverPckg(elevNet.ComsChan)
	go comsManager.SendPckgs(elevNet.ComsChan)
	go elevNet.ListenToBroadcast(elevNet.ComsChan)
	//go coms.SendPckgToAll(coms.ComsChan)
	
	//msg:=message.ConstructPckg("129.241.187.153","connectTo", "test")
	//for i:=0;i<1;i++{
	//	network.NetChan.SendUDP<-msg
	//}

	<-c
	
	
}
