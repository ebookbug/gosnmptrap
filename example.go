// example
package gosnmptrap

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Hello World!")
	socket,err := net.ListenUDP("udp4",&net.UDPAddr{
		IP:net.IPv4(0,0,0,0),
		Port:162,
	})
	if err !=nil{
		panic(err)
	}
	defer socket.Close()
}

func HandleUdp(data []byte){
	trap,err := ParseUdp(data)
	if err !=nil{
		fmt.Println("Err",err.Error())
	}
	fmt.Println(trap.Version,trap.Community,trap.EnterpriseId,trap.Address)
	for k,v :=range trap.Values{
		fmt.Printf("%s = %s\n",k,v);
	}
}