package network

import "coms"
import "fmt"



func DeliverPckg(pckgChan coms.ComsChannels){
    msg:=make([]byte,100)
    for{
		fmt.Println("in deliverPckg")
        msg=<-pckgChan.RecvPckg
        pckg :=Bytestrm2pckg(msg)
    
        switch pckg.Msg_type{
        case "PING":
            fmt.Println("The msg is of type PING")
		case "test":
			fmt.Println("hei")
        default:
            fmt.Println("not able to read msg header")
        }
    }
}
