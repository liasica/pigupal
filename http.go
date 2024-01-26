// Copyright (C) liasica. 2024-present.
//
// Created at 2024-01-25
// Based on pigupal by liasica, magicrolan@qq.com.

package pigupal

import (
	"html/template"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	_ = indexTmpl.Execute(w, map[string]string{
		"hi":   "Hi~ ˘¿˘",
		"ip":   CurrentIP,
		"port": CurrentPort,
	})
}

var indexTmpl *template.Template

func StartHttpServer() {
	indexTmpl, _ = template.New("index").Parse(TemplateIndex)

	http.HandleFunc("/", index)

	err := http.ListenAndServe(":56002", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
