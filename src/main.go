package main

import "runtime"
import "coms"



func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) 
	
	test :=coms.ConstructPckg("129.241.187.255","test", "cake or death?")

	go coms.SendPckgTo("129.241.187.255","20011",test)
	
	go coms.ListenToBroadcast("129.241.187.255","20011")
	
	for true{
	}
}
