package network

import "coms"
import "fmt"



func DeliverPckg(pckgChan coms.ComsChannels){
    fmt.Println("inni deliver")
    msg:=make([]byte,100)
    for{
        msg=<-pckgChan.RecvPckg
        pckg :=coms.Bytestrm2pckg(msg)
    
        switch pckg.Msg_type{
        case "PING":
            fmt.Println("The msg is of type PING")
        default:
            fmt.Println("not able to read msg header")
        }
    }
}
