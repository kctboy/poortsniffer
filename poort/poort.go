package poort

import (
	"net"
	"strconv"
	"time"
)

type ScanResult struct {
	port  int
	state string
}

func Scanport(protocol, hostname string, port int) ScanResult {
	result := ScanResult{port: port}
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		result.state = "closed"
		return result
	}
	defer conn.Close()
	result.state = "open"
	return result
}

func InitialScan(hostname string) []ScanResult {
	var results []ScanResult

	for i := 1; i <= 1024; i++ {
		results = append(results, Scanport("tcp", hostname, i))
	}
	return results

}
