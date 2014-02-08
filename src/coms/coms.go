
package coms

//FYYYYY
var MY_IP = GetMyIP()


const TARGET_PORT = "20011"
const LISTEN_PORT = "30011"

type ComsChannels struct{
    RecvPckg chan Msg_pckg
    SendPckg chan Msg_pckg  
	SendBcast chan Msg_pckg

}    
var ComsChan ComsChannels

func ComsChanInit(){
    ComsChan.RecvPckg = make(chan Msg_pckg,255)
    ComsChan.SendPckg = make(chan Msg_pckg,255)
	ComsChan.SendBcast = make(chan Msg_pckg,255)
}
