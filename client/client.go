package client

import (
	"fmt"
	"net/rpc"
	"strings"
	"strconv"
	// "errors"
    "github.com/jordansavant/elevators.go/server"
)

type Client struct {
	rpcClient *rpc.Client
}

func New() *Client {
	// create the work and his schedule and start him working
	c := Client{}
	return &c
}

func (c *Client) Start(addr string) {
	fmt.Println("starting client")

	// Establish the connection to the adddress of the
	// RPC server
	rpcClient, err := rpc.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	c.rpcClient = rpcClient
}

func (c *Client) AddWorker(name string, schedule string) {

	// convert schedule string to list
	// eg: 2:4_5:1_3:5_1:0
	pairs := strings.Split(schedule, "_")
	sched := make([]server.WorkerSchedulePair, len(pairs))
	for i, pair := range pairs {
		s := strings.Split(pair, ":")
		sched[i].Floor, _ = strconv.Atoi(s[0])
		sched[i].Seconds, _ = strconv.Atoi(s[1])
	}

	// add worker to server
	wresp := server.WorkerResponse{}
	err := c.rpcClient.Call("Server.AddWorker", &server.WorkerRequest{Name: name, Schedule: sched},  &wresp)
	if err != nil {
		fmt.Println("Server Error:", err)
		panic(err)
	}
	fmt.Println(wresp.Message, schedule)
}

func (c *Client) GetSnapshot() *server.SnapshotResponse {
	resp := server.SnapshotResponse{}
	err := c.rpcClient.Call("Server.GetSnapshot", &server.SnapshotRequest{}, &resp)
	if err != nil {
		fmt.Println("Server Error:", err)
		panic(err)
	}
	return &resp
}

func (c *Client) End() {
	fmt.Println("closing client connection")
	c.rpcClient.Close()
}