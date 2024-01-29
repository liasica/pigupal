// Copyright (C) liasica. 2024-present.
//
// Created at 2024-01-29
// Based on pigupal by liasica, magicrolan@qq.com.

package pigupal

import "os"

var CFG *Config

type Config struct {
	IpServer     string
	RconServer   string
	RconPassword string
	BackupDir    string
}

func Initialize() {
	CFG = &Config{
		IpServer:     os.Getenv("IP_SERVER"),
		RconServer:   os.Getenv("RCON_SERVER"),
		RconPassword: os.Getenv("RCON_PASSWORD"),
		BackupDir:    os.Getenv("BACKUP_DIR"),
	}
}
