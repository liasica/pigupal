// Copyright (C) liasica. 2024-present.
//
// Created at 2024-01-26
// Based on pigupal by liasica, magicrolan@qq.com.

package rcon

import (
	"log"
	"math/rand"
	"net"
	"time"
)

type Rcon struct {
	address  string
	password []byte
	client   *Client

	// options
	keepalive bool
	deadline  time.Duration
	timeout   time.Duration
}

// New Rcon
// https://tech.palworldgame.com/server-commands
// https://developer.valvesoftware.com/wiki/Source_RCON_Protocol
func New(address string, password string, options ...Option) *Rcon {
	r := &Rcon{
		address:   address,
		password:  []byte(password),
		keepalive: false,
		deadline:  time.Second,
		timeout:   time.Second * 3,
	}
	for _, option := range options {
		option(r)
	}
	return r
}

func (r *Rcon) Connect() (err error) {
	var conn net.Conn
	conn, err = net.DialTimeout("tcp", r.address, r.timeout)
	if err != nil {
		return
	}

	r.client = NewClient(conn)

	go r.client.Traffic()

	// auth
	err = r.client.Send(PacketIDAuth, PacketTypeAuth, r.password)
	if err != nil {
		return
	}

	for {
		select {
		case p := <-r.client.packet:
			go r.handlePacket(p)
		case err = <-r.client.exit:
			return
		}
	}
}

func (r *Rcon) handlePacket(p *Packet) {
	// log.Printf("%#v", p)
	if p.id == PacketIDAuth {
		log.Println("login success")
		err := r.client.Send(PackedID(rand.Int31()), PacketTypeExecCommandOrResponse, CommandShowPlayers)
		if err != nil {
			log.Println(err)
		}
	}
}
