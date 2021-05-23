package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"fr.serli.f1/application/packets"
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
 * Testing parsing only once...
 */
func TestInterfaceParsing(t *testing.T) {
	var packetHeader, decodedHeader packets.PacketHeader
	packetHeader = packets.PacketHeader{
		PacketId:                1,
		PlayerCarIndex:          0,
		PacketFormat:            2020,
		GameMajorVersion:        1,
		GameMinorVersion:        18,
		FrameIdentifier:         60234,
		SessionUID:              1234567,
		SessionTime:             200.0,
		SecondaryPlayerCarIndex: 255,
		PacketVersion:           1,
	}
	buff := new(bytes.Buffer)

	err := byteWriter(buff)(&packetHeader)

	assert.Nil(t, err)

	errReader := byteReader(buff)(&decodedHeader)

	assert.Nil(t, errReader)
	assert.Equal(t, packetHeader, decodedHeader)
}

func byteWriter(buff *bytes.Buffer) func(data interface{}) error {
	return func(data interface{}) error {
		err := binary.Write(buff, binary.LittleEndian, data)
		if err != nil {
			fmt.Printf("Some error  %v", err)
			return errors.New("cannot write binary")
		}
		return nil
	}
}
