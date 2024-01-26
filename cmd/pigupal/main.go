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
)

func main() {
	// reading env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Can not find .env file", err)
	}

	if strings.ToUpper(os.Getenv("DYNAMIC_IP")) == "TRUE" {
		go pigupal.StartIpClient()
	}

	go pigupal.StartHttpServer()

	select {}
}
