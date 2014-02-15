package coms
import("net")


var TcpConsMap = map[string]net.Conn{
}


const TARGET_PORT = "20011"
const LISTEN_PORT = "30011"

type ComsChannels struct{
	RecvPckg chan []byte
	SendPckg chan []byte  
	RecvBcast chan []byte
	SendBcast chan []byte

}    
var ComsChan ComsChannels

func ComsChanInit(){
	ComsChan.RecvPckg = make(chan []byte,255)
	ComsChan.SendPckg = make(chan []byte,255)
	ComsChan.SendBcast = make(chan []byte,255)
}

type TcpChannels struct{
	connect_to chan bool
	dead_elev chan bool
	new_conn chan net.Conn
}
	
var tcpChan TcpChannels

func tcpChanInit(){
	tcpChan.connect_to = make(chan bool)
	tcpChan.dead_elev = make(chan bool)
	tcpChan.new_conn = make(chan net.Conn)
}
