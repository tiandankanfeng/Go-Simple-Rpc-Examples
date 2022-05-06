package main

import (
   "fmt"
   "log"
   "net/rpc"
)

type RpcRequest struct {
   A, B int
}

type RpcResponse struct {
	Res int
}

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Synchronous call
	args := &RpcRequest{7,8}
	var reply RpcResponse
	err = client.Call("mathUtils.Multiply", args, &reply)
	if err != nil {
		log.Fatal("MathUtil error:", err)
	}
	fmt.Printf("MathUtil: %d*%d=%d", args.A, args.B, reply)

}