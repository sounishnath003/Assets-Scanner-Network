package main

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/sounishnath003/network-scanner/junk"
	"github.com/sounishnath003/network-scanner/worker"
)

var activeIps []string
var wg sync.WaitGroup

func main() {

	host := "192.168.0."

	for i := 1; i < 10; i++ {
		wg.Add(1)
		ip := host + strconv.Itoa(i)

		go fn0(ip)
	}

	fmt.Println("Workers: Waiting for workers to finish...")
	wg.Wait()
	fmt.Println("active ips", activeIps)
	fmt.Println("execution of program completed!!")
}

func fn0(ip string) {
	var activeIp string
	var err error

	wg.Add(1)
	defer wg.Done()

	go func() {
		activeIp, err = worker.Fn3(ip, &wg)
		if err != nil {
			fmt.Println(err)
		} else {
			activeIps = append(activeIps, activeIp)
			wg.Add(1)
			go junk.Junker(activeIp, &wg)
		}
	}()
}
