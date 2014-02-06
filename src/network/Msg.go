package Network

import "coms"



func DeliverPckg(coms.ComsChan.RecvPckg){
    msg:=make([]byte)
    msg<-coms.ComsChan.RecvPckg
    pckg :=coms.Bytestrm2pckg(msg)
    
    switch pckg.msg_type{
    case PING:
        fmt.Println("The msg is of type PING")
    default:
        fmt.Println("not able to read msg header")
    }   
}
