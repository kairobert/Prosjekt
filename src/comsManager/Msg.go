package comsManager

import "elevNet"
import "fmt"
import "message"




func DeliverMsg(networkChan elevNet.ElevNetChannels){
    for{
        msg:=<-networkChan.RecvMsg
    
        switch msg.Msg_type{
        case "connectTo":
            fmt.Println("The msg is of type udp")
			networkChan.ConnectToElev<-msg.From
		case "test":
			fmt.Println("tcp msg recieved")
        default:
            fmt.Println("not able to read msg header")
        }
    }
}

func SendMsg(msg message.Message,sendChan elevNet.ElevNetChannels){ //TTEST
	for{
		select{
		case <-NetChan.SendUDP:
			sendChan.SendBcast<-msg
		}
	}
}


