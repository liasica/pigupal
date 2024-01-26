// Copyright (C) liasica. 2024-present.
//
// Created at 2024-01-26
// Based on pigupal by liasica, magicrolan@qq.com.

package rcon

import (
	"bytes"
	"encoding/binary"
	"errors"
)

type PacketType int32

func (typ PacketType) Valid() bool {
	switch typ {
	case PacketTypeResponseValue, PacketTypeExecCommandOrResponse, PacketTypeAuth:
		return true
	default:
		return false
	}
}

const (
	PacketTypeResponseValue         PacketType = 0 // SERVERDATA_RESPONSE_VALUE
	PacketTypeExecCommandOrResponse PacketType = 2 // SERVERDATA_EXECCOMMAND
	PacketTypeAuth                  PacketType = 3 // SERVERDATA_AUTH
)

type PackedID int32

func (id PackedID) Valid() bool {
	switch id {
	case PacketIDAuth:
		return true
	default:
		return false
	}
}

const (
	PacketIDAuth PackedID = iota
)

const (
	packetEmptySize  = 2
	packetHeaderSize = 4
	packetIDSize     = 4
	packetTypeSize   = 4
)

var packetEmpty = []byte{0x00, 0x00}

type Packet struct {
	id   PackedID
	typ  PacketType
	body []byte
}

// NewPacket create packet
func NewPacket(id PackedID, typ PacketType, body []byte) (p *Packet) {
	p = &Packet{
		id:   id,
		typ:  typ,
		body: body,
	}
	return p
}

// GetBody Getting packet body, remove EMPTY
func (p *Packet) GetBody() []byte {
	return p.body
}

// Encode RCON packet data
// Size:    int32 <4 Bytes little endian>
// ID:      int32 <4 Bytes little endian>
// Type:    int32 <4 Bytes little endian>
// Body:    []byte (at least 1 Byte)
// End:     0x00 (1 Byte)
func (p *Packet) Encode() []byte {
	size := packetIDSize + packetTypeSize + len(p.body) + packetEmptySize

	buffer := bytes.NewBuffer(make([]byte, 0, packetHeaderSize+size))

	// put header [SIZE]
	_ = binary.Write(buffer, binary.LittleEndian, uint32(size))

	// put [ID]
	_ = binary.Write(buffer, binary.LittleEndian, p.id)

	// put [TYPE]
	_ = binary.Write(buffer, binary.LittleEndian, p.typ)

	// put [BODY]
	buffer.Write(p.body)

	// put [EMPTY]
	buffer.Write(packetEmpty)

	return buffer.Bytes()
}

// Decode RCON packet data
func (p *Packet) Decode(data []byte) error {
	// read [ID]
	r := bytes.NewReader(data)
	err := binary.Read(r, binary.LittleEndian, &p.id)
	if err != nil {
		return errors.New("read packet id error: " + err.Error())
	}

	// read [TYPE]
	err = binary.Read(r, binary.LittleEndian, &p.typ)
	if err != nil {
		return errors.New("read packet type error: " + err.Error())
	}

	if p.typ == PacketTypeExecCommandOrResponse && p.id == -1 {
		return ErrorAuthFailed
	}

	// Determine whether the [TYPE] is valid
	if !p.typ.Valid() {
		return errors.New("packet type is unknown")
	}

	// read [BODY] and remove all [EMPTY]
	p.body = bytes.TrimFunc(data[packetIDSize+packetTypeSize:], func(r rune) bool {
		return r == 0
	})

	return nil
}
