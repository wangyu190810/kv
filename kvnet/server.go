package kvnet

import (
	"fmt"
	"kv/config"
	"net"
	"strconv"
	"time"
)

func Runserver(address string) {
	ln, err := net.Listen("tcp", address)
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			//handle error
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		content := ""
		nums, err := conn.Read(buf)
		//fmt.Println(nums)
		if nums == 0 {
			fmt.Print("client lost")
			break
		} else if err != nil {
			fmt.Print("client err %v", nil)
		} else {
			if nums < 1024 {
				content = content + string(buf[:nums])
				fmt.Print(content)
				config.Data.Command(content)
				content = ""
			} else {
				content = content + string(buf[:nums])
			}
			buf = make([]byte, 1024)
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte(daytime))

		}

	}
}
