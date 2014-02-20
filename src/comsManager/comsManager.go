package comsManager
import("message")
  


type NetChannels struct{
	SendUDP chan message.Message

}    
var NetChan NetChannels

func NetChanInit(){
	NetChan.SendUDP= make(chan message.Message,255)
}
//func InitElevList()[]Elev{

//    ElevList:=make([]Elev,999)
//}

// Mail ElevPackage Delivery Message 
