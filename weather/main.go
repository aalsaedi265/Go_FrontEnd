package main

import (
	
	"log"
	"weather/p2p"
)

func main(){
	tr := p2p.NewTCPTransport(":3000")

	log.Fatal(tr.ListenAndAccept())

	select{}
}