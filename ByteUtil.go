package common

import (
	"sync"
	"encoding/base64"
	"encoding/binary"
)

type byteUtil struct {
	mutex sync.Mutex
}

// Default big endian
var Byte = byteUtil{}

func (this byteUtil) BytesToUint64(b []byte, littleEndian... bool) (v uint64) {
	if littleEndian != nil && len(littleEndian) > 0 && littleEndian[0] {
		v = binary.LittleEndian.Uint64(b)
	} else {
		v = binary.BigEndian.Uint64(b)
	}
	return
}

func (this byteUtil) BytesToUint32(b []byte, littleEndian... bool) (v uint32) {
	if littleEndian != nil && len(littleEndian) > 0 && littleEndian[0] {
		v = binary.LittleEndian.Uint32(b)
	} else {
		v = binary.BigEndian.Uint32(b)
	}
	return
}

func (this byteUtil) BytesToUint16(b []byte, littleEndian... bool) (v uint16) {
	if littleEndian != nil && len(littleEndian) > 0 && littleEndian[0] {
		v = binary.LittleEndian.Uint16(b)
	} else {
		v = binary.BigEndian.Uint16(b)
	}
	return
}

func (this byteUtil) Uint64toBytes(b []byte, v uint64, littleEndian... bool) {
	if littleEndian != nil && len(littleEndian) > 0 && littleEndian[0] {
		binary.LittleEndian.PutUint64(b, v)
	} else {
		binary.BigEndian.PutUint64(b, v)
	}
}

func (this byteUtil) Uint32toBytes(b []byte, v uint32, littleEndian... bool) {
	if littleEndian != nil && len(littleEndian) > 0 && littleEndian[0] {
		binary.LittleEndian.PutUint32(b, v)
	} else {
		binary.BigEndian.PutUint32(b, v)
	}
}

func (this byteUtil) Uint16toBytes(b []byte, v uint16, littleEndian... bool) {
	if littleEndian != nil && len(littleEndian) > 0 && littleEndian[0] {
		binary.LittleEndian.PutUint16(b, v)
	} else {
		binary.BigEndian.PutUint16(b, v)
	}
}

func (this byteUtil) Uint8toBytes(b []byte, v uint8) {
	b[0] = byte(v)
}

func (this byteUtil) BytesToUint8(b []byte) (v uint8) {
	return uint8(b[0])
}

func (this byteUtil) Base64Encode(value []byte) []byte {
	encoded := make([]byte, base64.URLEncoding.EncodedLen(len(value)))
	base64.URLEncoding.Encode(encoded, value)
	return encoded
}

func (this byteUtil) Base64Decode(value []byte) ([]byte, error) {
	decoded := make([]byte, base64.URLEncoding.DecodedLen(len(value)))
	b, err := base64.URLEncoding.Decode(decoded, value)
	if err != nil {
		return nil, err
	}
	return decoded[:b], nil
}
