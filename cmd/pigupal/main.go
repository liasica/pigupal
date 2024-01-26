// Copyright (C) liasica. 2024-present.
//
// Created at 2024-01-25
// Based on pigupal by liasica, magicrolan@qq.com.

package main

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"

	"pigupal"
	"pigupal/pkg/rcon"
)

func main() {
	// reading env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("can not find .env file: %v", err)
	}

	if strings.ToUpper(os.Getenv("DYNAMIC_IP")) == "TRUE" {
		go pigupal.StartIpClient()
	}

	go pigupal.StartHttpServer()

	r := rcon.New(os.Getenv("RCON_SERVER"), os.Getenv("RCON_PASSWORD"))
	err = r.Connect()
	if err != nil {
		log.Fatalf("connect rcon server failed: %v", err)
	}

	select {}
}
