# Golang implement for SNMP Trap

gosnmptrap is an open-source SNMP Trap Server library for Go(Golang).

This allow you setup a Trap Server to receive traps from your device such as Cisco switch.

By now,gosnmptrap suppport

1. Snmp V1 Trap
2. Snmp V2 Trap

## install

The easiest way to install is via go get:
	
    go get github.com/ebookbug/gosnmptrap
	
## example

	package main
	
	import (
		"fmt"
		"net"
		"github.com/ebookbug/gosnmptrap"
	)
	
	func main() {
		fmt.Println("Start a new UDPServr")	
		socket,err := net.ListenUDP("udp4",&net.UDPAddr{
			IP:net.IPv4(0,0,0,0),
			Port:162,
		})
		if err !=nil{
			panic(err)
		}
		defer socket.Close()
		
		for{
			buf := make([]byte,2048)
			read,from,_:=socket.ReadFromUDP(buf)
			fmt.Println("Get msg from ",from.IP)
			go HandleUdp(buf[:read])
		}
	}
	
	func HandleUdp(data []byte){
		trap,err := gosnmptrap.ParseUdp(data)
		if err !=nil{
			fmt.Println("Err",err.Error())
		}
		fmt.Println(trap.Version,trap.Community,trap.EnterpriseId,trap.Address)
		for k,v :=range trap.Values{
			fmt.Printf("%s = %s\n",k,v);
		}
	}
	
## license

Apache 2.0 licence