package main

import (
	"fmt"
	"sync"

	"github.com/sounishnath003/network-scanner/host"
	"github.com/sounishnath003/network-scanner/worker"
)

var activeIps []string
var wg sync.WaitGroup

func main() {

	ips, err := host.Hosts("192.168.0.1/24")

	if err != nil {
		panic(err)
	}

	for _, ip := range ips {
		wg.Add(1)
		go worker.IpParser(ip, &wg, &activeIps)
	}

	fmt.Println("Workers: Waiting for workers to pull it up...")
	wg.Wait()
	fmt.Println("active ips", activeIps)
	fmt.Println("execution of program completed!!")
}