// Copyright (C) liasica. 2024-present.
//
// Created at 2024-01-25
// Based on pigupal by liasica, magicrolan@qq.com.

package pigupal

import (
	"errors"
	"fmt"
	"html/template"
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

func ip(w http.ResponseWriter, r *http.Request) {
	CurrentIP, _ = getIP(r)
	_, _ = fmt.Fprintf(w, "CurrentIP = %s", CurrentIP)
}

func index(w http.ResponseWriter, r *http.Request) {
	// _, _ = fmt.Fprintf(w, "当前服务器地址为: %s:%s\n或使用域名: nas.liasica.com:%s", CurrentIP, CurrentPort, CurrentPort)
	_ = indexTmpl.Execute(w, map[string]string{
		"hi":   "Hi~ ˘¿˘",
		"ip":   CurrentIP,
		"port": CurrentPort,
	})
}

var indexTmpl *template.Template

func StartHttpServer() {
	indexTmpl, _ = template.New("index").Parse(TemplateIndex)
	// err = t.ExecuteTemplate(out, "T", "<script>alert('you have been pwned')</script>")

	http.HandleFunc("/", index)
	http.HandleFunc("/ip", ip)

	err := http.ListenAndServe(":19900", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
