// Copyright (C) liasica. 2024-present.
//
// Created at 2024-01-26
// Based on pigupal by liasica, magicrolan@qq.com.

package pigupal

import (
	"io"
	"log"
	"net/http"
	"time"
)

func getIp(addr string) {
	res, err := http.Get(addr)
	if err != nil {
		log.Println(err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	var b []byte
	b, err = io.ReadAll(res.Body)

	if err != nil {
		log.Printf("Get IP failed: %v", err)
	} else {
		CurrentIP = string(b)
	}
}

func StartIpClient(addr string) {
	log.Println("using [IP_SERVER]:", addr)

	ticker := time.NewTicker(time.Second * 5)
	for range ticker.C {
		getIp(addr)
	}
}
