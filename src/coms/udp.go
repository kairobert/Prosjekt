package coms

import (
	"net"

	
)

const UDP_PORT ="20000"//All Elevs listen to this Broadcast Port


func SendPckgToAll(pckgChan ComsChannels){
    bcastIP:=GetBroadcastIP(GetMyIP())
    
	serverAddr, err := net.ResolveUDPAddr("udp",bcastIP+":"+UDP_PORT)
	if err != nil {return}

	con, err := net.DialUDP("udp", nil, serverAddr)	
	if err != nil {return}
	
	for {
		msg:=make([]byte,100)
		msg=<-pckgChan.SendBcast
		con.Write(msg)
	}		
}

func ListenToBroadcast(pckgChan ComsChannels) {
	myIp :=GetMyIP()
	bcastIP:=GetBroadcastIP(myIp)
	

	serverAddr, err := net.ResolveUDPAddr("udp",bcastIP+":"+UDP_PORT)
	if err != nil { return }
	
	psock, err := net.ListenUDP("udp4", serverAddr)	
	if err != nil { return }
	
	buf := make([]byte,255)
 
  	for {
  	    _, remoteAddr, err := psock.ReadFromUDP(buf[0:])
    		if err != nil { return }
		
    		if remoteAddr.IP.String() != myIp {
    	    		pckgChan.RecvPckg<-buf
    		}       
      }	
}

