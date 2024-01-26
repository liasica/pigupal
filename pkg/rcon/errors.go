// Copyright (C) liasica. 2024-present.
//
// Created at 2024-01-26
// Based on pigupal by liasica, magicrolan@qq.com.

package rcon

import "errors"

var (
	// ErrorIncompletePacket The data packet is incomplete and needs to wait to receive the data packet.
	ErrorIncompletePacket = errors.New("incomplete packet")
	ErrorAuthFailed       = errors.New("auth failed")
)
