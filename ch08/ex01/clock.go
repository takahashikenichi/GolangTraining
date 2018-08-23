// clock1は時刻を定期的に書き出すTCPサーバです
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var port = flag.Int("port", 8080, "listen port")

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // 例：接続が切れた
			continue
		}
		go handleConn(conn) //  一度に一つの接続を処理する
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // 例：クライアントとの接続が切れた
		}
		time.Sleep(1 * time.Second)
	}
}
