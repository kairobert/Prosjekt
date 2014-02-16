package coms
import("net")


var TcpConsMap = map[string]net.Conn{
	}



const TARGET_PORT = "20011"
const LISTEN_PORT = "30011"

type ComsChannels struct{
	RecvPckg chan []byte
	SendPckg chan []byte  
	SendBcast chan []byte
	ConnectToElev chan string
}
    
var ComsChan ComsChannels

func ComsChanInit(){
	ComsChan.RecvPckg = make(chan []byte,255)
	ComsChan.SendPckg = make(chan []byte,255)
	ComsChan.SendBcast = make(chan []byte,255)
	ComsChan.ConnectToElev = make(chan string,255)
}

type TcpChannels struct{
	connect_to chan bool
	dead_elev chan bool
	new_conn chan net.Conn
}
	
var tcpChan TcpChannels

func TcpChanInit(){
	tcpChan.connect_to = make(chan bool, 255)
	tcpChan.dead_elev = make(chan bool, 255)
	tcpChan.new_conn = make(chan net.Conn, 255)
}
