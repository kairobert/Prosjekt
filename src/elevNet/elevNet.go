package elevNet
import("net")
import("message")



var TcpConsMap = map[string]net.Conn{
	}



const TARGET_PORT = "20011"
const LISTEN_PORT = "30011"

type ElevNetChannels struct{
	RecvMsg chan message.Message
	SendMsg chan message.Message  
	SendBcast chan message.Message
	ConnectToElev chan string
}
    
var ElevNetChan ElevNetChannels

func ElevNetChanInit(){
	ElevNetChan.RecvMsg = make(chan message.Message,255)
	ElevNetChan.SendMsg = make(chan message.Message,255)
	ElevNetChan.SendBcast = make(chan message.Message,255)
	ElevNetChan.ConnectToElev = make(chan string,255)

}

type TcpChannels struct{
	connect_to chan bool
	dead_elev chan bool
	new_conn chan net.Conn
	send_msg chan message.Message
}
	
var tcpChan TcpChannels

func TcpChanInit(){
	tcpChan.connect_to = make(chan bool, 255)
	tcpChan.dead_elev = make(chan bool, 255)
	tcpChan.new_conn = make(chan net.Conn, 255)
	tcpChan.send_msg = make(chan message.Message)
}
