// Copyright (C) liasica. 2024-present.
//
// Created at 2024-01-29
// Based on pigupal by liasica, magicrolan@qq.com.

package pigupal

import (
	"fmt"
	"os"
	"sort"
	"time"
)

type Backup struct {
	Name string
	Size string
	time time.Time
}

func LastBackups() (list []*Backup) {
	entries, err := os.ReadDir(CFG.BackupDir)
	if err != nil {
		return
	}

	for _, e := range entries {
		info, _ := e.Info()
		mt := info.ModTime()
		name := mt.Format("2006-01-02 15:04:05")
		list = append(list, &Backup{
			Name: name,
			Size: fmt.Sprintf("%.2f MB", float64(info.Size())/1024/1024),
			time: mt,
		})
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].time.After(list[j].time)
	})

	right := 5
	if len(list) < 5 {
		right = len(list)
	}

	return list[:right]
}
