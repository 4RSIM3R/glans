package main

import (
	"net"
	"time"

	"github.com/go-mysql-org/go-mysql/client"
	"github.com/go-mysql-org/go-mysql/server"
	_ "github.com/go-sql-driver/mysql"
	"github.com/siddontang/go-log/log"
	"polinema.ac.id/walek/src"
)

func main() {

	pool, err := client.NewPoolWithOptions(
		"127.0.0.1:3306",
		"root", "",
		"kolam",
		client.WithPoolLimits(15, 150, 30),
		client.WithLogFunc(log.Debugf),
		client.WithNewPoolPingTimeout(5*time.Second),
	)

	if err != nil {
		log.Fatal(err)
	}

	l, err := net.Listen("tcp", ":4000")

	if err != nil {
		log.Fatal(err)
	}

	// Accept a new connection once
	c, err := l.Accept()
	if err != nil {
		log.Fatal(err)
	}

	conn, err := server.NewConn(c, "root", "", &src.CustomHandler{
		Pool: pool,
	})

	if err != nil {
		log.Fatal(err)
	}

	// as long as the client keeps sending commands, keep handling them
	for {
		if err := conn.HandleCommand(); err != nil {
			log.Fatal(err)
		}
	}

}
