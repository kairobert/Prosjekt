package network

import "coms"
import "fmt"




func DeliverPckg(networkChan coms.ComsChannels){
    msg:=make([]byte,100)
    for{
        msg=<-networkChan.RecvPckg
        pckg :=Bytestrm2pckg(msg)
    
        switch pckg.Msg_type{
        case "connectTo":
            fmt.Println("The msg is of type udp")
			networkChan.ConnectToElev<-pckg.From
		case "test":
			fmt.Println("tcp msg recieved")
        default:
            fmt.Println("not able to read msg header")
        }
    }
}

func SendPckgs(sendChan coms.ComsChannels){ //TTEST
	for{
		select{
		case p:=<-NetChan.SendUDP:
			msg:=Pckg2bstream(p)
			sendChan.SendBcast<-msg
		}
	}
}


