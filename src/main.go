package main

import "runtime"
import "coms"
import "fmt"

const BCAST_IP = "129.241.187.255"
const LISTEN_PORT = "20011"
const TARGET_PORT = "20022"


func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) 
	
	test :=coms.ConstructPckg("129.241.187.255","test", "cake or death?")

	go coms.SendPckgTo(BCAST_IP,TARGET_PORT ,test)
	
	go coms.ListenToBroadcast(BCAST_IP,LISTEN_PORT )
	
	fmt.Println("hallaa")
	for true{
	}
}
