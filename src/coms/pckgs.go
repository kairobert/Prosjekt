package coms

import "strings"

type Msg_pckg struct{
	from string
	msg_type string
	payload string
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

func ConstructPckg(adr string, typ string, msg string) Msg_pckg{
	return Msg_pckg{adr,typ,msg}
}
