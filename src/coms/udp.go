package coms

import (
	"net"
	"fmt"
)

func SendPckgTo(ipAdr string, port string, p Msg_pckg){
	serverAddr, err := net.ResolveUDPAddr("udp",ipAdr+":"+port)
	con, err := net.DialUDP("udp", nil, serverAddr)
	
	if err != nil {
		fmt.Println("fail")
	}
	
	bstream := Pckg2bstream(p)
	
	for i:=0; i<10; i++ {
		
		con.Write(bstream)
	}		
}

func ListenToBroadcast(ipAdr string, port string){
	serverAddr, err := net.ResolveUDPAddr("udp",ipAdr+":"+port)
	psock, err := net.ListenUDP("udp4", serverAddr)
	
	if err != nil { return }
	buf := make([]byte,1024)
 
  	for {    		
    		if err != nil { return }
    		_, remoteAddr, _ := psock.ReadFromUDP(buf)
    		if remoteAddr.IP.String() != MY_IP {
    			fmt.Printf("oida")
		fmt.Printf("%s\n",buf)
		}
	 }	
		
}


