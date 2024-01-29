// Copyright (C) liasica. 2024-present.
//
// Created at 2024-01-25
// Based on pigupal by liasica, magicrolan@qq.com.

package pigupal

import (
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

type IndexData struct {
	Hi         string
	Ip         string
	Port       string
	Backups    []*Backup
	BackupsLen int
}

func index(w http.ResponseWriter, r *http.Request) {
	backups := LastBackups()
	_ = indexTmpl.Execute(w, &IndexData{
		Hi:         "Hi~ ˘¿˘",
		Ip:         CurrentIP,
		Port:       CurrentPort,
		Backups:    backups,
		BackupsLen: len(backups),
	})
}

var indexTmpl *template.Template

func StartHttpServer() {
	indexTmpl, _ = template.New("index").Parse(TemplateIndex)

	http.HandleFunc("/", index)
	imagesFs, _ := fs.Sub(Images, "assets/images")
	http.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.FS(imagesFs))))

	err := http.ListenAndServe(":56002", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
