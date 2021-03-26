package port

import (
	"net"
	"strconv"
	"time"
)

type ScanResult struct {
	Port  int
	State string
}

func ScanPort(protocol, hostname string, port int) ScanResult {
	result := ScanResult{Port: port}

	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 5*time.Second)

	if err != nil {
		result.State = "CLOSE"
		return result
	}

	defer conn.Close()
	result.State = "OPEN"
	return result
}

func InitialScan(hostname string) []ScanResult {
	var results []ScanResult

	for i := 0; i < 20; i++ {
		results = append(results, ScanPort("tcp", hostname, i))
	}

	return results
}
