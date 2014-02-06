package main

import "runtime"
import "coms"
import "fmt"
import "network"

const BCAST_IP = "129.241.187.255"
const LISTEN_PORT = "20022"
const TARGET_PORT = "20011"


func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) 
	sleepChan :=make(chan int)
	//test :=coms.ConstructPckg("129.241.187.255","PING", "cake or death?")
	coms.ComsChanInit()

	//go coms.SendPckgTo(BCAST_IP,TARGET_PORT ,test)
	
	go coms.ListenToBroadcast(BCAST_IP, LISTEN_PORT,coms.ComsChan )
	go network.DeliverPckg(coms.ComsChan)
	fmt.Println("hallaa")
	
	<-sleepChan
}
