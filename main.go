package main

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/sounishnath003/network-scanner/worker"
)

var activeIps []string
var wg sync.WaitGroup

func main() {

	host := "192.168.0."

	for i := 1; i < 10; i++ {
		wg.Add(1)
		ip := host + strconv.Itoa(i)

		go worker.IpParser(ip, &wg, &activeIps)
	}

	fmt.Println("Workers: Waiting for workers to finish...")
	wg.Wait()
	fmt.Println("active ips", activeIps)
	fmt.Println("execution of program completed!!")
}
