package main

import (
	"adnpa/id-generater/api/pb"
	_ "adnpa/id-generater/init"
	"adnpa/id-generater/internal/service"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

func main() {
	var port = flag.Int("port", 10000, "The server port")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterIdServer(s, &service.IdService{})

	// TODO 服务注册&健康检查

	log.Printf("listen on %d\n", *port)
	go func() {
		err = s.Serve(lis)
		if err != nil {
			panic("failed to start grpc:" + err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	// TODO 服务注销
}
