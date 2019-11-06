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
	rpcClient, _ := rpc.Dial("tcp", addr)
	c.rpcClient = rpcClient
}

func (c *Client) AddWorker(name string, schedule string) {

	// convert schedule string to list
	// eg: 2:4 5:1 3:5 1:0
	pairs := strings.Split(schedule, " ")
	fmt.Println(schedule, pairs)
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
	fmt.Println(wresp.Message)
}

func (c *Client) End() {
	c.rpcClient.Close()
}
