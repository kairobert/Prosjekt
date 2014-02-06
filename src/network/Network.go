package Network 
import "coms"
import "fmt"










type Elev struct{
    ip string
    sendPort string
    listenPort string
}

//func InitElevList()[]Elev{

//    ElevList:=make([]Elev,999)
//}

func DeliverPckg(){
    for {
        temp :="" 
        temp=<-coms.St_chan
        fmt.Println(temp)
    }   
}
