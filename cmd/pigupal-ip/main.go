// Copyright (C) liasica. 2024-present.
//
// Created at 2024-01-26
// Based on pigupal by liasica, magicrolan@qq.com.

package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

// getIP returns the ip address from the http request
func getIP(r *http.Request) (string, error) {
	ips := r.Header.Get("X-Forwarded-For")
	splitIps := strings.Split(ips, ",")

	if len(splitIps) > 0 {
		// get last IP in list since ELB prepends other user defined IPs, meaning the last one is the actual client IP.
		netIP := net.ParseIP(splitIps[len(splitIps)-1])
		if netIP != nil {
			return netIP.String(), nil
		}
	}

	found, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	netIP := net.ParseIP(found)
	if netIP != nil {
		found = netIP.String()
		if found == "::1" {
			return "127.0.0.1", nil
		}
		return found, nil
	}

	return "", errors.New("IP not found")
}

func main() {
	http.HandleFunc("/ip", func(writer http.ResponseWriter, request *http.Request) {
		found, _ := getIP(request)
		log.Println("获取到IP:", found)
		_, _ = fmt.Fprintf(writer, found)
	})

	err := http.ListenAndServe(":56001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
