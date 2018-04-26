package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

// TimeServer ...
type TimeServer int64

// Args ...
type Args struct{}

// GiveServerTime ...
func (t *TimeServer) GiveServerTime(args *Args, reply *int64) error {
	*reply = time.Now().Unix()
	log.Printf("%d", *reply)
	return nil
}

func main() {
	timeserver := new(TimeServer)
	rpc.Register(timeserver)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("Listen error :", e)
	}
	http.Serve(l, nil)
}
