package coms

import (
	"net"
	"fmt"
)


func SendPckgToAll(ipAdr string, port string, thisChan ComsChannels){
	serverAddr, err := net.ResolveUDPAddr("udp",ipAdr+":"+port)
	con, err := net.DialUDP("udp", nil, serverAddr)	
	if err != nil {
		fmt.Println("fail")
	}
	
	msg:=make([]byte)
	msg=<-RecvPckg
	con.Write(msg)
			
}

func ListenToBroadcast(ipAdr string, port string, Coms.Pckg_chan) {
	serverAddr, err := net.ResolveUDPAddr("udp",ipAdr+":"+port)
	if err != nil { return }
	
	psock, err := net.ListenUDP("udp4", serverAddr)	
	if err != nil { return }
	
	buf := make([]byte,255)
 
  	for {
  	    _, remoteAddr, err := psock.ReadFromUDP(buf)
    	if err != nil { return }
    	if remoteAddr.IP.String() != MY_IP {
    	    Pckg_chan<-buf
    	}       
    }	
}

