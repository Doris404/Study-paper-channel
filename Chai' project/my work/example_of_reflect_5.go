package main

import "proto"
import "fmt"

func main(){
	pb := proto.Message{}
	proto.Clone(pb)
	fmt.Println("Done")
}