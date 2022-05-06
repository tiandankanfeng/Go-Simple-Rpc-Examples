package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)


type RpcRequest struct {
	A, B int
}

type RpcResponse struct {
	Res int
}

type MathUtil struct {

}

// declared itf - this is the standard declaration of Rpc Service In GoLang
func (t *MathUtil) Multiply(args *RpcRequest, reply *RpcResponse) error {
	reply.Res = args.A * args.B
	return nil
}

func main() {

	mathUtil := new(MathUtil)
	// 别名姿势
	err := rpc.RegisterName("mathUtils", mathUtil)
	if err != nil {
		return
	}

	// 绑定协议 - 后续基于此协议进行数据的传输
	rpc.HandleHTTP()

	// 监听端口
	l, e := net.Listen("tcp", ":1234")

	if e != nil {
		log.Fatal("listen error:", e)
	}
	fmt.Println("server is working")

	// 开启服务
	http.Serve(l, nil)

}

