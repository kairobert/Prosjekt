package coms
import(
	"net"
	"fmt"	
	"strings"	
)

const CON_ATMPTS = 10
const TCP_PORT = "30000" //All elevators will listen to this port for TCP connections



func HandleTCPCom(){	
	go listenTcpCon()

	for {
	fmt.Println("in select")
		select{
		case newTcpCon:=<- tcpChan.new_conn:
			handleNewCon(newTcpCon)
		}//end select
	}//end for
}


func listenMsg(con net.Conn){
	fmt.Println("inside listen Msg")
	msg := make([]byte,1024)
	addr := con.RemoteAddr()
	fmt.Println(addr)
    for {
		_, err := con.Read(msg[0:])
	    if err!=nil {
			//fmt.Println("error in listen")			
		}else{
			ComsChan.RecvPckg<-msg
			fmt.Println("sendt msg on channel to network")
		}
	}
}

func listenTcpCon(){
	localAddr, err := net.ResolveTCPAddr("tcp",":"+TCP_PORT)
	sock, err := net.ListenTCP("tcp", localAddr)
	if err != nil { return }

 
	for{
		con, err := sock.Accept()
			if err != nil {
				return
			}else{
   				fmt.Println("connection ok")
   				tcpChan.new_conn<-con
				fmt.Println("sendt con on chan")
   				//"send on new ip channel"
   			
   			}
   	}
}	

func sendTcpMsg(msg []byte, ipAddr string){
	con, ok :=TcpConsMap[ipAddr]
	if ok{
		_, err := con.Write(msg)
		if err!=nil{
			fmt.Println("failed to send msg")
		}
	}else{
		fmt.Println("error, not a connection")
	}
}


func ConnectTcp(ipAdr string){
	atmpts:=0
	for atmpts < CON_ATMPTS{
		serverAddr, err := net.ResolveTCPAddr("tcp",ipAdr+":"+TCP_PORT)
		if err != nil {
			fmt.Println("Error Resolving Address")
			atmpts++
		}else{
			con, err := net.DialTCP("tcp", nil,serverAddr);
			if err != nil {
				fmt.Println("Error DialingTCP")
				atmpts++
			}else{
				fmt.Println("connection ok")
				tcpChan.new_conn<-con
				break
			}
		}//end BIG if/else		
	}//end for
}

func handleNewCon(con net.Conn){
	fmt.Println("handle new Con")
	ip:= getConIp(con)

	_, ok := TcpConsMap[ip]
	if !ok{	
		fmt.Println("in if")
		TcpConsMap[ip]=con
		go listenMsg(con)
	}else{
		fmt.Println("connection already excist")
	}
}



func getConIp(con net.Conn)(ip string){
	split:=strings.Split(con.RemoteAddr().String(),":") //splits ip from port
	conIp :=split[0]
	return conIp
	
}
