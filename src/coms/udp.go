package coms

import (
	"net"

	
)


func SendPckgToAll( port string, pckgChan ComsChannels){
    bcastIP:=GetBroadcastIP(GetMyIP())
    
	serverAddr, err := net.ResolveUDPAddr("udp",bcastIP+":"+port)
	if err != nil {return}

	con, err := net.DialUDP("udp", nil, serverAddr)	
	if err != nil {return}
	
	
	msg:=make([]byte,100)
	msg=<-pckgChan.SendPckg
	con.Write(msg)
			
}

func ListenToBroadcast(port string, pckgChan ComsChannels) {
	bcastIP:=GetBroadcastIP(GetMyIP())

	serverAddr, err := net.ResolveUDPAddr("udp",bcastIP+":"+port)
	if err != nil { return }
	
	psock, err := net.ListenUDP("udp4", serverAddr)	
	if err != nil { return }
	
	buf := make([]byte,255)
 
  	for {
  	    _, remoteAddr, err := psock.ReadFromUDP(buf)
    	if err != nil { return }

    	if remoteAddr.IP.String() != MY_IP {
    	    pckgChan.RecvPckg<-buf
    	}       
    }	
}

