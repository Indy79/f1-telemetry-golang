package packets

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
 * Testing parsing only once...
 */
func TestHeaderParsing(t *testing.T) {

	var network bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.
	dec := gob.NewDecoder(&network) // Will read from network.
	header := PacketHeader{PacketId: 1, PlayerCarIndex: 0, PacketFormat: 2020, GameMajorVersion: 1, GameMinorVersion: 18}
	err := enc.Encode(header)
	assert.Nil(t, err)

	fmt.Println(network.Bytes())

	var decodedHeader PacketHeader
	err = dec.Decode(&decodedHeader)
	assert.Nil(t, err)
	assert.Equal(t, header, decodedHeader)
	assert.Equal(t, header.GameMinorVersion, decodedHeader.GameMinorVersion)
	assert.Equal(t, header.PacketVersion, decodedHeader.PacketVersion)
}
