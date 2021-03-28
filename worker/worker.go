package worker

import (
	"errors"
	"fmt"
	"net"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/sounishnath003/network-scanner/junk"
)

var activeIp []string

func Fn0(ip string, wg *sync.WaitGroup) {

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

func Fn1(ip string, wg *sync.WaitGroup) {

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

func Fn3(ip string, wg *sync.WaitGroup) (string, error) {

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

func IpParser(ip string, wg *sync.WaitGroup, activeIps *[]string) {
	var activeIp string
	var err error

	wg.Add(1)
	defer wg.Done()

	go func() {
		activeIp, err = Fn3(ip, wg)
		if err != nil {
			fmt.Println(err)
		} else {
			*activeIps = append(*activeIps, activeIp)
			wg.Add(1)
			go junk.Junker(activeIp, wg)
		}
	}()
}