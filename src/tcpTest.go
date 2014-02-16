package main

import "runtime"
import "coms"
import "network"
//import "fmt"
//import "strings"



const TARGET_IP = "129.241.187.142"



func main() {
	coms.TcpChanInit()
	coms.ComsChanInit()
	network.NetChanInit()
	c := make(chan int)
	runtime.GOMAXPROCS(runtime.NumCPU())
	
	 
	go coms.HandleTCPCom()
	go coms.ListenToBroadcast(coms.ComsChan)
	go network.DeliverPckg(coms.ComsChan)
	go network.SendPckgs(coms.ComsChan)
	go coms.ListenToBroadcast(coms.ComsChan)
	//go coms.SendPckgToAll(coms.ComsChan)
	
	//msg:=network.ConstructPckg("129.241.187.152","connectTo", "test")
	//for i:=0;i<1;i++{
	//	network.NetChan.SendUDP<-msg
	//}

	<-c
	
	
}
