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

func GetMsg(){
    for {
        temp :="" 
        temp=<-coms.St_chan
        fmt.Println(temp)
    }   
}
