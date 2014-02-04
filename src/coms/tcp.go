package coms


import "fmt"
import "net"


func ConnectToElev(ipAdr string, port string, msg Msg_pckg){
   serverAddr, err := net.ResolveTCPAddr("tcp",ipAdr+":"+port)
	con, err := net.DialTCP("tcp", nil, serverAddr);
	
	if err != nil{
		fmt.Println("fail")
	}
   
   
   bstream :=Pckg2bstream(msg)   
   con.Write(bstream)

}


         
