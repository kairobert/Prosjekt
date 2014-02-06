package coms

import "strings"

const MSG_PCKG_SIZE = 1024

type Msg_pckg struct{
	from string
	msg_type string //order, deadElev, auction, connect to me
	payload string
}

var Bcast_chan = make(chan Msg_pckg, MSG_PCKG_SIZE)
var St_chan =make(chan string,100)


func ConstructPckg(adr string, typ string, msg string) Msg_pckg{
	return Msg_pckg{adr,typ,msg}
}

//not generic, could use reflect...
func Pckg2bstream(p Msg_pckg) []byte{
	msg := p.from +"~"+ p.msg_type +"~"+ p.payload
	return []byte(msg+"\x00")
}

//not generic, could use reflect..
func Bytestrm2pckg(p []byte) Msg_pckg{
	msg_string := string(p[:])
	msg_array := strings.Split(msg_string, "~")
	pckg := Msg_pckg{msg_array[0], msg_array[1], msg_array[2]}
	return pckg
}

