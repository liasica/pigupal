// Copyright (C) liasica. 2024-present.
//
// Created at 2024-01-26
// Based on pigupal by liasica, magicrolan@qq.com.

package rcon

import (
	"bufio"
	"encoding/binary"
	"errors"
	"log"
	"net"
)

type Client struct {
	net.Conn

	packet chan *Packet
	exit   chan error
}

func NewClient(conn net.Conn) *Client {
	return &Client{
		Conn:   conn,
		packet: make(chan *Packet),
	}
}

func (c *Client) Send(id PackedID, typ PacketType, body []byte) (err error) {
	b := NewPacket(id, typ, body).Encode()
	_, err = c.Conn.Write(b)
	return
}

func (c *Client) Traffic() {
	for {
		b, err := c.Read()
		if err != nil {
			if errors.Is(err, ErrorIncompletePacket) {
				continue
			}
			c.exit <- err
			return
		}

		p := new(Packet)
		err = p.Decode(b)
		if err != nil {
			c.exit <- err
			return
		}

		c.packet <- p
	}
}

// func (c *Client) Read() (b []byte, err error) {
// 	b = make([]byte, 4096)
// 	var n int
// 	n, err = c.Conn.Read(b)
// 	if n == 0 {
// 		return nil, ErrorIncompletePacket
// 	}
// 	log.Printf("%s", b)
// 	return b[packetHeaderSize:], nil
// }

func (c *Client) Read() ([]byte, error) {
	r := bufio.NewReader(c.Conn)
	header, _ := r.Peek(packetHeaderSize)
	if len(header) < 4 {
		return nil, ErrorIncompletePacket
	}

	size := int(binary.LittleEndian.Uint32(header))
	total := size + packetHeaderSize
	buffered := r.Buffered()
	if buffered < total {
		x, _ := r.Peek(buffered)
		log.Printf("%s", x)
		// log.Println(r.Discard(buffered))
		return nil, ErrorIncompletePacket
	}

	b, err := r.Peek(total)
	if err != nil {
		return nil, err
	}

	// _, err = r.Discard(total)
	// if err != nil {
	// 	return nil, err
	// }

	return b[packetHeaderSize:], nil
}
