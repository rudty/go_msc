package main

import "net"

type client struct {
	ClientID uint32
	Conn     net.Conn
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
