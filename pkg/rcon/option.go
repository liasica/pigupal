// Copyright (C) liasica. 2024-present.
//
// Created at 2024-01-26
// Based on pigupal by liasica, magicrolan@qq.com.

package rcon

import "time"

type Option func(r *Rcon)

func WithKeepAlive(keepalive bool) Option {
	return func(r *Rcon) {
		r.keepalive = keepalive
	}
}

func WithDialTimeout(timeout time.Duration) Option {
	return func(r *Rcon) {
		r.timeout = timeout
	}
}

func WithDeadline(timeout time.Duration) Option {
	return func(r *Rcon) {
		r.deadline = timeout
	}
}
