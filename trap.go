// trap
package gosnmptrap

import (
	"fmt"
	"github.com/cdevr/WapSNMP"
)

type Trap struct{
	Version int16
	Community string
	GeneralTrap int32
	SpeicalTrap int32
	EnterpriseId string
	Values map[string]interface{}
}

func ParseUdp(data []byte) (Trap,error){
	trap := Trap{}
	seq,err := wapsnmp.DecodeSequence(data)
	if err !=nil{
		fmt.Println(err)
	}
//	for i:=0;i<len(seq);i++{
//		fmt.Println(seq[i])
//	}
	
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
		fmt.Printf("OID: %s\n",vrsp[1])
		fmt.Printf("Agent Address: %s\n",vrsp[2])
		fmt.Printf("Generic Trap: %d\n",vrsp[3])
		fmt.Printf("Special Trap: %d\n",vrsp[4])
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
	fmt.Printf("\n");
	
	return trap,nil
}