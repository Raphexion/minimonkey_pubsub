package main

// Encode a message with code and payload
func Encode(code byte, payload []byte) []byte {
	size := uint16(len(payload))
	var high, low uint8 = uint8(size >> 8), uint8(size & 0xff)
	header := []byte{code, low, high}
	return append(header, payload...)
}
