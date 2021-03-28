package main

import (
	"errors"
	"fmt"
	"net"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/sounishnath003/network-scanner/junk"
)

var activeIp []string
var wg sync.WaitGroup

func main() {

	host := "192.168.0."

	for i := 1; i < 255; i++ {
		wg.Add(1)
		ip := host + strconv.Itoa(i)

		var activeIp string
		var err error

		go func() {
			activeIp, err = fn3(ip)
			if err != nil {
				fmt.Println(err)
			} else {
				wg.Add(1)
				go junk.Junker(activeIp, &wg)
			}
		}()

	}

	fmt.Println("Workers: Waiting for workers to finish...")
	wg.Wait()
	fmt.Println("active ips", activeIp)
	fmt.Println("execution of program completed!!")
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

func fn3(ip string) (string, error) {

	defer wg.Done()

	bytes, err := exec.Command("powershell", "ping", ip).CombinedOutput()

	if err != nil {
	} else {
		if strings.Contains(string(bytes), "Destination host unreachable") {
			fmt.Println("BAD HOST: " + ip + " host is unreachable...")
			return ip, errors.New("BAD HOST: " + ip + " host is unreachable...")
		}
	}
	activeIp = append(activeIp, ip)
	fmt.Println("running on ip:", ip, "succeed...")
	return ip, nil

}
package main

import (
	"errors"
	"fmt"
	"net"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/sounishnath003/network-scanner/junk"
)

var activeIp []string
var wg sync.WaitGroup

func main() {

	host := "192.168.0."

	for i := 1; i < 255; i++ {
		wg.Add(1)
		ip := host + strconv.Itoa(i)

		var activeIp string
		var err error

		go func() {
			activeIp, err = fn3(ip)
			if err != nil {
				fmt.Println(err)
			} else {
				wg.Add(1)
				go junk.Junker(activeIp, &wg)
			}
		}()

	}

	fmt.Println("Workers: Waiting for workers to finish...")
	wg.Wait()
	fmt.Println("active ips", activeIp)
	fmt.Println("execution of program completed!!")
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

func fn3(ip string) (string, error) {

	defer wg.Done()

	bytes, err := exec.Command("powershell", "ping", ip).CombinedOutput()

	if err != nil {
	} else {
		if strings.Contains(string(bytes), "Destination host unreachable") {
			fmt.Println("BAD HOST: " + ip + " host is unreachable...")
			return ip, errors.New("BAD HOST: " + ip + " host is unreachable...")
		}
	}
	activeIp = append(activeIp, ip)
	fmt.Println("running on ip:", ip, "succeed...")
	return ip, nil

}
