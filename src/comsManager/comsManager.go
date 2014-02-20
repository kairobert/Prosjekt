package comsManager
 



type NetChannels struct{
	SendUDP chan Msg_pckg

}    
var NetChan NetChannels

func NetChanInit(){
	NetChan.SendUDP= make(chan Msg_pckg,255)
}
//func InitElevList()[]Elev{

//    ElevList:=make([]Elev,999)
//}


