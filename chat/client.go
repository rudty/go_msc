package main

import (
	"errors"
	"io"
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

func (c *client) ReadMessageInto(buf []byte) (int, error) {
	var header [4]byte
	_, err := c.Conn.Read(header[:])
	if err != nil {
		return 0, err
	}

	length := int(header[0])
	length |= int(header[1]) << 8
	length |= int(header[2]) << 16
	length |= int(header[3]) << 24

	readLength := 0

	for readLength < length {
		l, err := c.Conn.Read(buf[:length])
		if err != nil {
			if err == io.EOF {
				return 0, errors.New("packet body read fail")
			}
			return 0, err
		}
		readLength += l
	}
	return length, nil
}
