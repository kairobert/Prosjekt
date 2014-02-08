package coms

import "net"



func SendToElev(ipAdr string, port string, pckg[]byte){
    serverAddr, err := net.ResolveTCPAddr("tcp",ipAdr+":"+port)
	if err != nil {return}

	con, err := net.DialTCP("tcp", nil, serverAddr);
	if err != nil {return}
	
	msg :=make([]byte,255)
	
	   
	con.Write(bstream)

}


         
