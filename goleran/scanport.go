package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)
func main() {
	start := time.Now()
	var wg sync.WaitGroup
	for i := 21; i < 65535; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("192.168.10.219:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				fmt.Printf("%s 关闭\n", address)
				return
			}
			conn.Close()
			fmt.Printf("%s 打开\n", address)
		}(i)
	}
	wg.Wait()
	elapsed := time.Since(start) / 1e9
	fmt.Printf("\n\n %d seconds", elapsed)

}

