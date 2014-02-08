
package coms

//FYYYYY
var MY_IP = GetMyIP()


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
