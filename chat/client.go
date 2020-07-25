package main

import (
	"errors"
	"io"
	"net"
	"sync/atomic"
)

var errorPacketBodyReadFail = errors.New("packet body read fail")
var errorPacketHeaderReadFail = errors.New("packet header read fail")

type client struct {
	ClientID uint32
	Conn     net.Conn
	UserData interface{}
}

var uniqueID uint32 = 0

// newClient 새로운 클라이언트를 만듭니다.
// clientID는 유니크하게 증가합니다
func newClient(conn net.Conn) *client {
	clientID := atomic.AddUint32(&uniqueID, 1)
	return &client{
		Conn:     conn,
		ClientID: clientID,
	}
}

// WriteMessage write int32 + ByteArray
func (c *client) WriteMessage(buf []byte) error {
	var header [4]byte
	length := len(buf)
	header[0] = byte(length)
	header[1] = byte(length >> 8)
	header[2] = byte(length >> 16)
	header[3] = byte(length >> 24)

	if err := c.WriteByteArray(header[:]); err != nil {
		return err
	}

	return c.WriteByteArray(buf[:])
}

func (c *client) WriteByteArray(buf []byte) error {
	sendIndex := 0
	for sendIndex < len(buf) {
		sendLength, err := c.Conn.Write(buf[sendIndex:])
		if err != nil {
			return err
		}

		sendIndex += sendLength
	}
	return nil
}

// readPacket 소켓에서 인자로 입력받은 buf에 꽉 찰때까지 계속 읽습니다.
func (c *client) readPacketInto(buf []byte) error {
	readLength := 0
	for readLength < len(buf) {
		l, err := c.Conn.Read(buf[readLength:])
		if err != nil {
			return err
		}
		readLength += l
	}
	return nil
}

// ReadMessageInto (int32 + byte array) 형식의 패킷을 읽습니다.
// 일단 int32를 읽고 그 길이만큼 byte array를 읽습니다.
// 구조 상 첫번째로 읽는 int32는 못읽고 연결이 끊어질 수 있지만
// 두번째로 읽는 byte array는 반드시 읽어야합니다.
func (c *client) ReadMessageInto(buf []byte) (int, error) {
	if err := c.readPacketInto(buf[:4]); err != nil {
		if err == io.EOF {
			return 0, err
		}
		return 0, errorPacketHeaderReadFail
	}

	length := int(buf[0])
	length |= int(buf[1]) << 8
	length |= int(buf[2]) << 16
	length |= int(buf[3]) << 24

	if err := c.readPacketInto(buf[:length]); err != nil {
		return 0, errorPacketBodyReadFail
	}

	return length, nil
}
