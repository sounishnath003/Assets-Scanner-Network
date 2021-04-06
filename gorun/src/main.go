package main

import (
	"fmt"
	"net"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"
)

var activeIp []string
var wg sync.WaitGroup

func main() {

	host := "192.168.0."

	for i := 1; i < 100; i++ {
		wg.Add(1)
		ip := host + strconv.Itoa(i)

		go fn1(ip)

	}

	fmt.Println("Workers: Waiting for workers to finish...")
	wg.Wait()
	fmt.Println("active ips", activeIp)

}

func fn0(ip string) {

	defer wg.Done()

	bytes, err := exec.Command("powershell", "ping", ip).CombinedOutput()

	if err != nil {
	} else {
		if strings.Contains(string(bytes), "Destination host unreachable") {
			fmt.Println("BAD HOST: " + ip + " host is unreachable...")
		} else {
			activeIp = append(activeIp, ip)
			fmt.Println("running on ip:", ip, "succeed...")
		}
	}

}

func fn1(ip string) {

	defer wg.Done()

	conn, err := net.DialTimeout("tcp", ip+":80", time.Second)

	if err != nil {
		fmt.Println("BAD HOST: " + ip + " host is unreachable...")
	} else {
		defer conn.Close()
		activeIp = append(activeIp, ip)
		fmt.Println("running on ip:", ip, "succeed...")
	}
}
