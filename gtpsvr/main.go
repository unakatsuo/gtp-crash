package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/wmnsk/go-gtp/gtpv1"
)

func main() {
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

	gtpAddr, err := net.ResolveUDPAddr("udp", "localhost:2563")
	if err != nil {
		panic(err)
	}
	gtpConn := gtpv1.NewUPlaneConn(gtpAddr)
	if err := gtpConn.EnableKernelGTP("gtp1", gtpv1.RoleGGSN); err != nil {
		gtpConn.Close()
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-sigCh
		gtpConn.Close()
		cancel()
	}()
	fmt.Println("Starting to listen and serve")
	if err := gtpConn.ListenAndServe(ctx); err != nil {
		log.Printf("Failed to listen on %s: %+v", gtpAddr, err)
	}

	//gtpConn.Close()
}
