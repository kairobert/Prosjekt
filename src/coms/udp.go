package coms

import (
	"net"
	"fmt"
)


func SendPckgToAll(ipAdr string, port string, p Msg_pckg){
	serverAddr, err := net.ResolveUDPAddr("udp",ipAdr+":"+port)
	con, err := net.DialUDP("udp", nil, serverAddr)
	
	if err != nil {
		fmt.Println("fail")
	}
	
	bstream := Pckg2bstream(p)
	con.Write(bstream)
			
}

func ListenToBroadcast(ipAdr string, port string, Pckg_chan) {
	serverAddr, err := net.ResolveUDPAddr("udp",ipAdr+":"+port)
	psock, err := net.ListenUDP("udp4", serverAddr)	
	if err != nil { return }
	
	buf := make([]byte,255)
 
  	for {    		
    		
    		_, remoteAddr, err := psock.ReadFromUDP(buf)
    		if err != nil { return }
    		if remoteAddr.IP.String() != MY_IP {
    		 
    		//fmt.Println(remoteAddr.IP.String())
    		//}
		    //fmt.Printf("%s\n",buf)
            
	    	
	 }	
		
}

