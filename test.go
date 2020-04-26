package main

import(
    "fmt"
)

type Block struct{
	bpm int
	name string

}
var blockchain []Block

func add(BPM int, Name string){
	var block Block
	block.bpm = BPM
	block.name = Name

	blockchain = append(blockchain, block);
}

func display(){
	len := len(blockchain)
	fmt.Print(len)
}

func main(){

	add(54, "thejas")
	display()
}

