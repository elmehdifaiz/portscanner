package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
	"os"
)

func scanner(server string) []int{
	var avalible_ports []int
	for i := 1; i <= 65535; i++{
		ip := server + ":" + strconv.Itoa(i)
		_, err := net.DialTimeout("tcp", ip, time.Duration(300)*time.Millisecond)
		if err == nil{
			avalible_ports = append(avalible_ports, i)
		}
	}
	return avalible_ports
}

func main() {
	fmt.Println("Cheking for available ports...")
    ports := scanner(os.Args[1])
    fmt.Println("Ports available: " ,ports)
}
