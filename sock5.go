package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

const (
	// AtypIPV4 IPv4
	AtypIPV4 = uint8(0x01)
	// AtypDomainName 域名
	AtypDomainName = uint8(0x03)
	// AtypIPV6 IPv6
	AtypIPV6 = uint8(0x04)
)

type Socks5Protocol struct {
	data []byte
	Cmd  uint8
	Atyp uint8
	Host string
	Port int32
}

func (sp *Socks5Protocol) ParseRequest() bool {
	dataLen := len(sp.data)
	if dataLen < 5 {
		return false
	}
	sp.Cmd = uint8(sp.data[1])
	sp.Atyp = uint8(sp.data[3])
	fmt.Println(sp.Atyp)

	switch sp.Atyp {
	case AtypIPV4:
		if dataLen < 10 {
			return false
		}
		sp.Host = net.IPv4(sp.data[4], sp.data[5], sp.data[6], sp.data[7]).String()
		sp.Port = int32(binary.BigEndian.Uint16(sp.data[8:10]))
	case AtypIPV6:
		if dataLen < 22 {
			return false
		}
		str := string(sp.data[4:20])
		fmt.Println(str)
		var ipByte net.IP = []byte(str) //sp.data[4:20]
		sp.Host = ipByte.String()
		sp.Port = int32(binary.BigEndian.Uint16(sp.data[20:22]))
	case AtypDomainName:
		hostLen := int(sp.data[4])
		if dataLen < (5 + hostLen + 2) {
			return false
		}
		sp.Host = string(sp.data[5 : 5+hostLen])
		sp.Port = int32(binary.BigEndian.Uint16(sp.data[5+hostLen : 5+hostLen+2]))
	default:
		return false
	}
	return true
}

func main() {

	//var buf = []byte{0x05, 0x01, 0x00, 0x01, 0x01, 0x02, 0x03, 0x04, 0xff, 0xff} // 1.2.3.4:65535
	var buf = []byte{0x05, 0x01, 0x00, 0x04, 0x61, 0x62, 0x63, 0x64, 0x61, 0x62, 0x63, 0x64, 0x61, 0x62, 0x63, 0x64, 0x61, 0x62, 0x63, 0x64, 0xff, 0xff}
	sp := &Socks5Protocol{data: buf}
	sp.ParseRequest()

	fmt.Println(sp)

}
