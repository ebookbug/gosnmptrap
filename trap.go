// trap
package gosnmptrap

import (
	"fmt"
	"net"
	"github.com/cdevr/WapSNMP"
)

type Trap struct{
	Address string
	Version int16
	Community string
	GeneralTrap int16
	SpeicalTrap int16
	EnterpriseId string
	Values map[string]interface{}
}

func ParseUdp(data []byte) (Trap,error){
	trap := Trap{}
	seq,err := wapsnmp.DecodeSequence(data)
	if err !=nil{
		fmt.Println(err)
		return trap,err
	}
	
	var community string
	snmpVer := seq[1].(int64)
	if snmpVer<=1{
		snmpVer++
	}
	trap.Version = int16(snmpVer)
	
	if snmpVer <3{
		community = seq[2].(string)
	}
	
	trap.Community = community
	
	var vrsp []interface{}
			
	switch seq[3].(type){
		case wapsnmp.UnsupportedBerType:
			vrsp,_ = wapsnmp.DecodeSequence(seq[3].(wapsnmp.UnsupportedBerType))
		default:
			vrsp = seq[3].([]interface{})
	}
		
	var varbinds []interface{}
	
	if snmpVer==1{
		trap.GeneralTrap = int16(vrsp[3].(int64))
		trap.SpeicalTrap = int16(vrsp[4].(int64))
		trap.EnterpriseId = vrsp[1].(wapsnmp.Oid).String()
		trap.Address = vrsp[2].(net.IP).String()
		varbinds = vrsp[6].([]interface{})
	}else{
		varbinds = vrsp[4].([]interface{})
	}
	
	trap.Values = make(map[string]interface{},len(varbinds))
	
	for i:=1;i<len(varbinds);i++ {
		varoid:= varbinds[i].([]interface{})[1].(wapsnmp.Oid)
		result := varbinds[i].([]interface{})[2]
		
		trap.Values[varoid.String()] = result
	}
	
	return trap,nil
}