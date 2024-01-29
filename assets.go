// Copyright (C) liasica. 2024-present.
//
// Created at 2024-01-25
// Based on pigupal by liasica, magicrolan@qq.com.

package pigupal

import (
	"embed"
	_ "embed"
)

var (
	//go:embed templates/index.html
	TemplateIndex string

	//go:embed assets/images/*
	Images embed.FS
)
