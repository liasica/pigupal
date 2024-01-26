// Copyright (C) liasica. 2024-present.
//
// Created at 2024-01-25
// Based on pigupal by liasica, magicrolan@qq.com.

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gorcon/rcon"
)

func main() {
	// pigupal.StartHttpServer()
	conn, err := rcon.Dial("10.10.10.20:25575", "XXXXXXXXXXX", func(s *rcon.Settings) {
		rcon.SetDeadline(time.Second * 60)
	})
	if err != nil {
		log.Fatal(err)
	}
	defer func(conn *rcon.Conn) {
		_ = conn.Close()
	}(conn)

	// go func() {
	// 	for {}
	// }()

	// packet := rcon.NewPacket(rcon.SERVERDATA_EXECCOMMAND, rcon.SERVERDATA_EXECCOMMAND_ID, "ShowPlayers")

	response, err := conn.Execute("ShowPlayers")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
}
