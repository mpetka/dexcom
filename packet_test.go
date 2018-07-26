package dexcom

import (
	"bytes"
	"testing"
)

func TestPacket(t *testing.T) {
	cases := []struct {
		cmd    Command
		params []byte
		msg    []byte
	}{
		{Ping, nil,
			[]byte{0x01, 0x06, 0x00, 0x0A, 0x5E, 0x65}},
		{ReadTransmitterID, nil,
			[]byte{0x01, 0x06, 0x00, 0x19, 0x0C, 0x47}},
		{ReadDatabasePageRange, []byte{0x04},
			[]byte{0x01, 0x07, 0x00, 0x10, 0x04, 0x8B, 0xB8}},
		{ReadDatabasePages, []byte{0x05, 0x26, 0x00, 0x00, 0x00, 0x01},
			[]byte{0x01, 0x0C, 0x00, 0x11, 0x05, 0x26, 0x00, 0x00, 0x00, 0x01, 0x5E, 0xC3}},
	}
	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			msg := marshalPacket(c.cmd, c.params)
			if !bytes.Equal(msg, c.msg) {
				t.Errorf("Command(%X, %X) == %X, want %X", c.cmd, c.params, msg, c.msg)
			}
		})
	}
}
