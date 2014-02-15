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
	c := make(chan int)
	runtime.GOMAXPROCS(runtime.NumCPU())
	 
	go coms.HandleTCPCom()
	go network.DeliverPckg(coms.ComsChan)
	
	

	<-c
	
	
}
