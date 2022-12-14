package utils

import (
	"bytes"
	"encoding/binary"
	"log"
)

// IntToBytes
func IntToBytes(num int64) []byte {

	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}
