package main

import(
	"bytes"
	"encoding/binary"

)

type CommandMeneger struct{
	Manager *MemoryManager
	IP *byte
	G bool
	
}

func createCommandMeneger()(*CommandMeneger){
	this:=new(CommandMeneger)
	this.Manager=createMemmoryManager()
	this.IP=this.Manager.CodeSegment
	this.G=false
	return this
}

func (c *CommandMeneger)loadProgramm(programm []byte){
	(*c).Manager.writeMemmory(CS,programm)
}

func (c *CommandMeneger)move(addr1 uint,addr2 uint){
	var Mem []byte=(*c).Manager.readMemmory(addr2,4)
	(*c).Manager.writeMemmory(addr1,Mem)

}

func (c *CommandMeneger)xor(addr1 uint,addr2 uint,addr3 uint){
	var Mem1 []byte=(*c).Manager.readMemmory(addr3,4)
	var Mem2 []byte=(*c).Manager.readMemmory(addr2,4)
	var result []byte=make([]byte, 0,4)
	for i:=range Mem1{
		var temp byte=Mem1[i]^Mem2[i]
		result=append(result,temp)
	}
	(*c).Manager.writeMemmory(addr1,result)
}

func (c *CommandMeneger)shl(addr1 uint,addr3 uint,count uint){
	var Mem1 []byte=(*c).Manager.readMemmory(addr3,4)
	var result []byte=make([]byte, 0,4)
	for i:=range Mem1{
		var temp byte=Mem1[i]<<count
		result=append(result,temp)
	}
	(*c).Manager.writeMemmory(addr1,result)
}

func (c *CommandMeneger)cmp(addr1 uint,addr2 uint){
	var Mem1 []byte=(*c).Manager.readMemmory(addr2,4)
	var Mem2 []byte=(*c).Manager.readMemmory(addr1,4)
	buf1:=bytes.NewBuffer(Mem1)
	buf2:=bytes.NewBuffer(Mem2)
	var op1 int=0 
	binary.Read(buf1, binary.LittleEndian, &op1)
	var op2 int=0  
	binary.Read(buf2, binary.LittleEndian, &op2)
	if op2>op1{
		(*c).G=true
		return
	}
	(*c).G=false
	
}

func (c *CommandMeneger)inc(addr1 uint){
	var Mem []byte=(*c).Manager.readMemmory(addr1,4)
	buf1:=bytes.NewBuffer(Mem)
	var op1 int=0
	binary.Read(buf1, binary.LittleEndian, &op1)
	op1++
	b:=make([]byte, 4)
	binary.BigEndian.PutUint32(b,uint32(op1))
	(*c).Manager.writeMemmory(addr1,b)
}

func (c *CommandMeneger)jump(addr1 uint){
	(*c).IP=&((*c).Manager.Memory[addr1])
}

func (c *CommandMeneger)add(addr1 uint,addr2 uint,addr3 uint){
	var Mem1 []byte=(*c).Manager.readMemmory(addr3,4)
	var Mem2 []byte=(*c).Manager.readMemmory(addr2,4)
	buf1:=bytes.NewBuffer(Mem1)
	buf2:=bytes.NewBuffer(Mem2)
	var op1 int=0 
	binary.Read(buf1, binary.LittleEndian, &op1)
	var op2 int=0  
	binary.Read(buf2, binary.LittleEndian, &op2)
	var result int=op1+op2
	b:=make([]byte, 4)
	binary.BigEndian.PutUint32(b,uint32(result))
	(*c).Manager.writeMemmory(addr1,b)
}