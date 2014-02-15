package coms
import(
	"net"
	"fmt"			
)

const CON_ATMPTS = 10
const TCP_PORT = "30000" //All elevators will listen to this port for TCP connections



func handleTCPCom(){
	go listenTcpMsg()
	go listenTcpCon()
	
	for {
		select{
		case <-tcpChan.dead_elev:
			fmt.Println("do something")
		case <-tcpChan.connect_to:
		default:
			fmt.Println("do something else")
		}//end select
	}//end for
}


//delete this later pls
func ConnectToElev(ipAdr string, port string, pckg []byte){
    serverAddr, err := net.ResolveTCPAddr("tcp",ipAdr+":"+port)
	if err != nil {return}

	con, err := net.DialTCP("tcp", nil, serverAddr);
	if err != nil {return}
	
		msg :=make([]byte,255)
		con.Write(msg)
}


func listenTcpMsg(){
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
   				fmt.Println("ok")
   				tcpChan.new_conn<-con
   				//"send on new ip channel"
   			
   			}
   	}
}	


func connectTcp(ipAdr string){
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

