package main

import (
	"net"
)

type client struct {
	ClientID uint32
	Conn     net.Conn
}

// WriteMessage write int32 + ByteArray
func (c *client) WriteMessage(buf []byte) error {
	var header [4]byte
	length := len(buf)
	header[0] = byte(length)
	header[1] = byte(length >> 8)
	header[2] = byte(length >> 16)
	header[3] = byte(length >> 24)

	if _, err := c.Conn.Write(header[:]); err != nil {
		return err
	}

	return c.WriteByteArray(buf[:])
}

func (c *client) WriteByteArray(buf []byte) error {
	var sendIndex int = 0
	for {
		sendLength, err := c.Conn.Write(buf[sendIndex:])
		if err != nil {
			return err
		}

		sendIndex += sendLength

		if sendLength >= len(buf) {
			return nil
		}
	}
}
