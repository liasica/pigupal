// Copyright (C) liasica. 2024-present.
//
// Created at 2024-01-26
// Based on pigupal by liasica, magicrolan@qq.com.

package rcon

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPacket(t *testing.T) {
	p := NewPacket(PacketIDAuth, PacketTypeAuth, []byte("TEST PASSWORD!"))
	encoded := p.Encode()
	t.Logf("Encoded: %v", encoded)

	decoded := new(Packet)
	err := decoded.Decode(encoded[4:])
	require.NoError(t, err)

	require.Equal(t, p.id, decoded.id)
	require.Equal(t, p.typ, decoded.typ)
	require.Equal(t, p.body, decoded.body)
}
