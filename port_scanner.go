package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports, result chan int) {
	for p := range ports {
		address := fmt.Sprintf("denobilisindri.in:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			result <- 0
			continue
		}
		conn.Close()
		result <- p
	}
}

func main() {
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}
	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()
	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}
	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("[+] Open Port : %d\n", port)
	}
}
