
package coms

//FYYYYY
var MY_IP = GetMyIP()

//const TARGET_IP =  getBroadcastIP()
const TARGET_PORT = "20011"
const LISTEN_PORT = "30011"

type ComsChannels struct
     Pckg_chan chan []byte
   
var ComsChan
}


var udp_con = make(chan Msg_pckg,255)


