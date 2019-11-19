package main
import (
	"fmt"
)
const (
	buffer=2097152
	Eax=4
	Ebx=8
	Ecx=12
	Esi=16
	DS=160
	CS=1184
)
type MemoryManager struct{
	Memory []byte
	Eax *byte
	Ebx	*byte
	Ecx	*byte
	Esi	*byte
	DataSegmanet *byte
	CodeSegment	*byte
	Labels map[string]uint
}

func createMemmoryManager()(*MemoryManager){
	this:=new(MemoryManager)
	this.Memory=make([]byte, buffer)
	this.Eax=&(this.Memory[4])
	this.Ebx=&(this.Memory[8])
	this.Ecx=&(this.Memory[12])
	this.Esi=&(this.Memory[16])
	this.DataSegmanet=&(this.Memory[160])
	this.CodeSegment=&(this.Memory[1184])
	this.Labels=make(map[string]uint,10)
	return this
}

func(p *MemoryManager) writeMemmory(addr uint,data []byte){
	for i:=range data{
		(*p).Memory[addr+uint(i)]=data[i]
	}
}

func (p *MemoryManager) readMemmory(addr uint, size uint)([]byte){
	return (*p).Memory[addr:addr+size]
}

func (p *MemoryManager) createLabel(name string,addr uint){
	(*p).Labels[name]=addr
}

func (p *MemoryManager)readLabel(name string)(res uint,e error){
	if key,ok:=(*p).Labels[name]; ok{
		return key,nil;
	}
		e=fmt.Errorf("Labels not found")
		return
	
}