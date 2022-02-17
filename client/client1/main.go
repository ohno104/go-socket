package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	Client()
}

func Client() {
	conn, err := net.Dial("tpc", "127.0.0.1:8088")
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		line = strings.Trim(line, "\r\n")

		if line == "exit" {
			fmt.Println("用戶退出客戶端")
			break
		}

		conent, err := conn.Write([]byte(line + "\n"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("客戶端發送了 %d 字節的數據到服務器\n", conent)
	}
}
