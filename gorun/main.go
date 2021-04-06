package main

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/sounishnath003/network-scanner/host"
	"github.com/sounishnath003/network-scanner/junk"
	"github.com/sounishnath003/network-scanner/worker"
)

var activeIps []string
var wg sync.WaitGroup
var mutex sync.RWMutex

func main() {
	runtime.GOMAXPROCS(100)
	ips, err := host.Hosts("192.168.0.5/24") // CIDR

	if err != nil {
		panic(err)
	}

	for _, ip := range ips {
		wg.Add(1)
		mutex.Lock()
		go worker.IpParser(ip, &wg, &activeIps, &mutex)
		mutex.Unlock()
	}

	fmt.Println("Workers: Waiting for workers to pull it up...")
	wg.Wait()
	fmt.Println("active ips", activeIps)

	mutex.Lock()
	junk.WritePayloadToFile()
	mutex.Unlock()

	fmt.Println("execution of program completed!!")
}
