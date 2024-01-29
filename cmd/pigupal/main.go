// Copyright (C) liasica. 2024-present.
//
// Created at 2024-01-25
// Based on pigupal by liasica, magicrolan@qq.com.

package main

import (
	"log"
	"strings"

	"github.com/joho/godotenv"

	"pigupal"
)

// palworld configs
// https://tech.palworldgame.com/optimize-game-balance
// https://pal-conf.bluefissure.com/
func main() {
	// reading env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("can not find .env file: %v", err)
	}

	pigupal.Initialize()

	if strings.ToUpper(pigupal.CFG.IpServer) != "" {
		go pigupal.StartIpClient(pigupal.CFG.IpServer)
	}

	go pigupal.StartHttpServer()

	// r := rcon.New(pigupal.CFG.RconServer, pigupal.CFG.RconPassword)
	// err = r.Connect()
	// if err != nil {
	// 	log.Fatalf("connect rcon server failed: %v", err)
	// }

	select {}
}
