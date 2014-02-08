package coms



import "net"


func ConnectToElev(ipAdr string, port string, msg Msg_pckg){
    serverAddr, err := net.ResolveTCPAddr("tcp",ipAdr+":"+port)
	if err != nil {return}

	con, err := net.DialTCP("tcp", nil, serverAddr);
	if err != nil {return}
	
	bstream :=Pckg2bstream(msg)   
	con.Write(bstream)

}


         
