// Copyright (C) liasica. 2024-present.
//
// Created at 2024-01-26
// Based on pigupal by liasica, magicrolan@qq.com.

package pigupal

type PackType int32

const (
	PackTypeResponseValue PackType = 0
	PackTypeExeccommand   PackType = 2
	PackTypeAuthResponse  PackType = 2
	PackTypeAuth          PackType = 3
)

// RCON data
// Size:    int32 <4 Bytes little endian>
// ID:      int32 <4 Bytes little endian>
// Type:    int32 <4 Bytes little endian>
// Body:    []byte (at least 1 Byte)
// End:     0x00 (1 Byte)
func encode() {
}

func decode() {
}
